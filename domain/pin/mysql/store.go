// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package pin

import (
	"fmt"
	"time"

	"github.com/documize/community/core/env"
	"github.com/documize/community/core/streamutil"
	"github.com/documize/community/domain"
	"github.com/documize/community/domain/store/mysql"
	"github.com/documize/community/model/pin"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Scope provides data access to MySQL.
type Scope struct {
	Runtime *env.Runtime
}

// Add saves pinned item.
func (s Scope) Add(ctx domain.RequestContext, pin pin.Pin) (err error) {
	row := s.Runtime.Db.QueryRow("SELECT max(sequence) FROM pin WHERE orgid=? AND userid=?", ctx.OrgID, ctx.UserID)
	var maxSeq int
	err = row.Scan(&maxSeq)

	if err != nil {
		maxSeq = 99
	}

	pin.Created = time.Now().UTC()
	pin.Revised = time.Now().UTC()
	pin.Sequence = maxSeq + 1

	stmt, err := ctx.Transaction.Preparex("INSERT INTO pin (refid, orgid, userid, labelid, documentid, pin, sequence, created, revised) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer streamutil.Close(stmt)

	if err != nil {
		err = errors.Wrap(err, "prepare pin insert")
		return
	}

	_, err = stmt.Exec(pin.RefID, pin.OrgID, pin.UserID, pin.FolderID, pin.DocumentID, pin.Pin, pin.Sequence, pin.Created, pin.Revised)
	if err != nil {
		err = errors.Wrap(err, "execute pin insert")
		return
	}

	return
}

// GetPin returns requested pinned item.
func (s Scope) GetPin(ctx domain.RequestContext, id string) (pin pin.Pin, err error) {
	stmt, err := s.Runtime.Db.Preparex("SELECT id, refid, orgid, userid, labelid as folderid, documentid, pin, sequence, created, revised FROM pin WHERE orgid=? AND refid=?")
	defer streamutil.Close(stmt)

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("prepare select for pin %s", id))
		return
	}

	err = stmt.Get(&pin, ctx.OrgID, id)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("execute select for pin %s", id))
		return
	}

	return
}

// GetUserPins returns pinned items for specified user.
func (s Scope) GetUserPins(ctx domain.RequestContext, userID string) (pins []pin.Pin, err error) {
	err = s.Runtime.Db.Select(&pins, "SELECT id, refid, orgid, userid, labelid as folderid, documentid, pin, sequence, created, revised FROM pin WHERE orgid=? AND userid=? ORDER BY sequence", ctx.OrgID, userID)

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("execute select pins for org %s and user %s", ctx.OrgID, userID))
		return
	}

	return
}

// UpdatePin updates existing pinned item.
func (s Scope) UpdatePin(ctx domain.RequestContext, pin pin.Pin) (err error) {
	pin.Revised = time.Now().UTC()

	var stmt *sqlx.NamedStmt
	stmt, err = ctx.Transaction.PrepareNamed("UPDATE pin SET labelid=:folderid, documentid=:documentid, pin=:pin, sequence=:sequence, revised=:revised WHERE orgid=:orgid AND refid=:refid")
	defer streamutil.Close(stmt)

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("prepare pin update %s", pin.RefID))
		return
	}

	_, err = stmt.Exec(&pin)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("execute pin update %s", pin.RefID))
		return
	}

	return
}

// UpdatePinSequence updates existing pinned item sequence number
func (s Scope) UpdatePinSequence(ctx domain.RequestContext, pinID string, sequence int) (err error) {
	stmt, err := ctx.Transaction.Preparex("UPDATE pin SET sequence=?, revised=? WHERE orgid=? AND userid=? AND refid=?")
	defer streamutil.Close(stmt)

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("prepare pin sequence update %s", pinID))
		return
	}

	_, err = stmt.Exec(sequence, time.Now().UTC(), ctx.OrgID, ctx.UserID, pinID)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("execute pin sequence update %s", pinID))
		return
	}

	return
}

// DeletePin removes folder from the store.
func (s Scope) DeletePin(ctx domain.RequestContext, id string) (rows int64, err error) {
	b := mysql.BaseQuery{}
	return b.DeleteConstrained(ctx.Transaction, "pin", ctx.OrgID, id)
}

// DeletePinnedSpace removes any pins for specified space.
func (s Scope) DeletePinnedSpace(ctx domain.RequestContext, spaceID string) (rows int64, err error) {
	b := mysql.BaseQuery{}
	return b.DeleteWhere(ctx.Transaction, fmt.Sprintf("DELETE FROM pin WHERE orgid=\"%s\" AND labelid=\"%s\"", ctx.OrgID, spaceID))
}

// DeletePinnedDocument removes any pins for specified document.
func (s Scope) DeletePinnedDocument(ctx domain.RequestContext, documentID string) (rows int64, err error) {
	b := mysql.BaseQuery{}
	return b.DeleteWhere(ctx.Transaction, fmt.Sprintf("DELETE FROM pin WHERE orgid=\"%s\" AND documentid=\"%s\"", ctx.OrgID, documentID))
}
