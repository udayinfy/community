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

// Package domain ...
package domain

import (
	"github.com/documize/community/model/account"
	"github.com/documize/community/model/activity"
	"github.com/documize/community/model/attachment"
	"github.com/documize/community/model/audit"
	"github.com/documize/community/model/block"
	"github.com/documize/community/model/category"
	"github.com/documize/community/model/doc"
	"github.com/documize/community/model/group"
	"github.com/documize/community/model/link"
	"github.com/documize/community/model/org"
	"github.com/documize/community/model/page"
	"github.com/documize/community/model/permission"
	"github.com/documize/community/model/pin"
	"github.com/documize/community/model/search"
	"github.com/documize/community/model/space"
	"github.com/documize/community/model/user"
)

// Store provides access to data store (database)
type Store struct {
	Account      AccountStorer
	Activity     ActivityStorer
	Attachment   AttachmentStorer
	Audit        AuditStorer
	Block        BlockStorer
	Category     CategoryStorer
	Document     DocumentStorer
	Group        GroupStorer
	Link         LinkStorer
	Organization OrganizationStorer
	Page         PageStorer
	Pin          PinStorer
	Permission   PermissionStorer
	Search       SearchStorer
	Setting      SettingStorer
	Space        SpaceStorer
	User         UserStorer
}

// SpaceStorer defines required methods for space management
type SpaceStorer interface {
	Add(ctx RequestContext, sp space.Space) (err error)
	Get(ctx RequestContext, id string) (sp space.Space, err error)
	PublicSpaces(ctx RequestContext, orgID string) (sp []space.Space, err error)
	GetViewable(ctx RequestContext) (sp []space.Space, err error)
	GetAll(ctx RequestContext) (sp []space.Space, err error)
	Update(ctx RequestContext, sp space.Space) (err error)
	Delete(ctx RequestContext, id string) (rows int64, err error)
}

// CategoryStorer defines required methods for category and category membership management
type CategoryStorer interface {
	Add(ctx RequestContext, c category.Category) (err error)
	Update(ctx RequestContext, c category.Category) (err error)
	Get(ctx RequestContext, id string) (c category.Category, err error)
	GetBySpace(ctx RequestContext, spaceID string) (c []category.Category, err error)
	GetAllBySpace(ctx RequestContext, spaceID string) (c []category.Category, err error)
	GetSpaceCategorySummary(ctx RequestContext, spaceID string) (c []category.SummaryModel, err error)
	Delete(ctx RequestContext, id string) (rows int64, err error)
	AssociateDocument(ctx RequestContext, m category.Member) (err error)
	DisassociateDocument(ctx RequestContext, categoryID, documentID string) (rows int64, err error)
	RemoveCategoryMembership(ctx RequestContext, categoryID string) (rows int64, err error)
	DeleteBySpace(ctx RequestContext, spaceID string) (rows int64, err error)
	GetDocumentCategoryMembership(ctx RequestContext, documentID string) (c []category.Category, err error)
	GetSpaceCategoryMembership(ctx RequestContext, spaceID string) (c []category.Member, err error)
	RemoveDocumentCategories(ctx RequestContext, documentID string) (rows int64, err error)
	RemoveSpaceCategoryMemberships(ctx RequestContext, spaceID string) (rows int64, err error)
}

// PermissionStorer defines required methods for space/document permission management
type PermissionStorer interface {
	AddPermission(ctx RequestContext, r permission.Permission) (err error)
	AddPermissions(ctx RequestContext, r permission.Permission, actions ...permission.Action) (err error)
	GetUserSpacePermissions(ctx RequestContext, spaceID string) (r []permission.Permission, err error)
	GetSpacePermissions(ctx RequestContext, spaceID string) (r []permission.Permission, err error)
	GetCategoryPermissions(ctx RequestContext, catID string) (r []permission.Permission, err error)
	GetCategoryUsers(ctx RequestContext, catID string) (u []user.User, err error)
	GetUserCategoryPermissions(ctx RequestContext, userID string) (r []permission.Permission, err error)
	GetUserDocumentPermissions(ctx RequestContext, documentID string) (r []permission.Permission, err error)
	GetDocumentPermissions(ctx RequestContext, documentID string) (r []permission.Permission, err error)
	DeleteDocumentPermissions(ctx RequestContext, documentID string) (rows int64, err error)
	DeleteSpacePermissions(ctx RequestContext, spaceID string) (rows int64, err error)
	DeleteUserSpacePermissions(ctx RequestContext, spaceID, userID string) (rows int64, err error)
	DeleteUserPermissions(ctx RequestContext, userID string) (rows int64, err error)
	DeleteCategoryPermissions(ctx RequestContext, categoryID string) (rows int64, err error)
	DeleteSpaceCategoryPermissions(ctx RequestContext, spaceID string) (rows int64, err error)
	DeleteGroupPermissions(ctx RequestContext, groupID string) (rows int64, err error)
}

// UserStorer defines required methods for user management
type UserStorer interface {
	Add(ctx RequestContext, u user.User) (err error)
	Get(ctx RequestContext, id string) (u user.User, err error)
	GetByDomain(ctx RequestContext, domain, email string) (u user.User, err error)
	GetByEmail(ctx RequestContext, email string) (u user.User, err error)
	GetByToken(ctx RequestContext, token string) (u user.User, err error)
	GetBySerial(ctx RequestContext, serial string) (u user.User, err error)
	GetActiveUsersForOrganization(ctx RequestContext) (u []user.User, err error)
	GetUsersForOrganization(ctx RequestContext, filter string) (u []user.User, err error)
	GetSpaceUsers(ctx RequestContext, spaceID string) (u []user.User, err error)
	GetUsersForSpaces(ctx RequestContext, spaces []string) (u []user.User, err error)
	UpdateUser(ctx RequestContext, u user.User) (err error)
	UpdateUserPassword(ctx RequestContext, userID, salt, password string) (err error)
	DeactiveUser(ctx RequestContext, userID string) (err error)
	ForgotUserPassword(ctx RequestContext, email, token string) (err error)
	CountActiveUsers() (c int)
	MatchUsers(ctx RequestContext, text string, maxMatches int) (u []user.User, err error)
}

// AccountStorer defines required methods for account management
type AccountStorer interface {
	Add(ctx RequestContext, account account.Account) (err error)
	GetUserAccount(ctx RequestContext, userID string) (account account.Account, err error)
	GetUserAccounts(ctx RequestContext, userID string) (t []account.Account, err error)
	GetAccountsByOrg(ctx RequestContext) (t []account.Account, err error)
	DeleteAccount(ctx RequestContext, ID string) (rows int64, err error)
	UpdateAccount(ctx RequestContext, account account.Account) (err error)
	HasOrgAccount(ctx RequestContext, orgID, userID string) bool
	CountOrgAccounts(ctx RequestContext) int
}

// OrganizationStorer defines required methods for organization management
type OrganizationStorer interface {
	AddOrganization(ctx RequestContext, org org.Organization) error
	GetOrganization(ctx RequestContext, id string) (org org.Organization, err error)
	GetOrganizationByDomain(subdomain string) (org org.Organization, err error)
	UpdateOrganization(ctx RequestContext, org org.Organization) (err error)
	DeleteOrganization(ctx RequestContext, orgID string) (rows int64, err error)
	RemoveOrganization(ctx RequestContext, orgID string) (err error)
	UpdateAuthConfig(ctx RequestContext, org org.Organization) (err error)
	CheckDomain(ctx RequestContext, domain string) string
}

// PinStorer defines required methods for pin management
type PinStorer interface {
	Add(ctx RequestContext, pin pin.Pin) (err error)
	GetPin(ctx RequestContext, id string) (pin pin.Pin, err error)
	GetUserPins(ctx RequestContext, userID string) (pins []pin.Pin, err error)
	UpdatePin(ctx RequestContext, pin pin.Pin) (err error)
	UpdatePinSequence(ctx RequestContext, pinID string, sequence int) (err error)
	DeletePin(ctx RequestContext, id string) (rows int64, err error)
	DeletePinnedSpace(ctx RequestContext, spaceID string) (rows int64, err error)
	DeletePinnedDocument(ctx RequestContext, documentID string) (rows int64, err error)
}

// AuditStorer defines required methods for audit trails
type AuditStorer interface {
	// Record logs audit entry using own DB Transaction
	Record(ctx RequestContext, t audit.EventType)
}

// DocumentStorer defines required methods for document handling
type DocumentStorer interface {
	Add(ctx RequestContext, document doc.Document) (err error)
	Get(ctx RequestContext, id string) (document doc.Document, err error)
	GetBySpace(ctx RequestContext, spaceID string) (documents []doc.Document, err error)
	TemplatesBySpace(ctx RequestContext, spaceID string) (documents []doc.Document, err error)
	DocumentMeta(ctx RequestContext, id string) (meta doc.DocumentMeta, err error)
	PublicDocuments(ctx RequestContext, orgID string) (documents []doc.SitemapDocument, err error)
	Update(ctx RequestContext, document doc.Document) (err error)
	UpdateGroup(ctx RequestContext, document doc.Document) (err error)
	ChangeDocumentSpace(ctx RequestContext, document, space string) (err error)
	MoveDocumentSpace(ctx RequestContext, id, move string) (err error)
	Delete(ctx RequestContext, documentID string) (rows int64, err error)
	DeleteBySpace(ctx RequestContext, spaceID string) (rows int64, err error)
	GetVersions(ctx RequestContext, groupID string) (v []doc.Version, err error)
}

// SettingStorer defines required methods for persisting global and user level settings
type SettingStorer interface {
	Get(area, path string) (val string, err error)
	Set(area, value string) error
	GetUser(orgID, userID, area, path string) (val string, err error)
	SetUser(orgID, userID, area, json string) error
}

// AttachmentStorer defines required methods for persisting document attachments
type AttachmentStorer interface {
	Add(ctx RequestContext, a attachment.Attachment) (err error)
	GetAttachment(ctx RequestContext, orgID, attachmentID string) (a attachment.Attachment, err error)
	GetAttachments(ctx RequestContext, docID string) (a []attachment.Attachment, err error)
	GetAttachmentsWithData(ctx RequestContext, docID string) (a []attachment.Attachment, err error)
	Delete(ctx RequestContext, id string) (rows int64, err error)
}

// LinkStorer defines required methods for persisting content links
type LinkStorer interface {
	Add(ctx RequestContext, l link.Link) (err error)
	SearchCandidates(ctx RequestContext, keywords string) (docs []link.Candidate, pages []link.Candidate, attachments []link.Candidate, err error)
	GetDocumentOutboundLinks(ctx RequestContext, documentID string) (links []link.Link, err error)
	GetPageLinks(ctx RequestContext, documentID, pageID string) (links []link.Link, err error)
	MarkOrphanDocumentLink(ctx RequestContext, documentID string) (err error)
	MarkOrphanPageLink(ctx RequestContext, pageID string) (err error)
	MarkOrphanAttachmentLink(ctx RequestContext, attachmentID string) (err error)
	DeleteSourcePageLinks(ctx RequestContext, pageID string) (rows int64, err error)
	DeleteSourceDocumentLinks(ctx RequestContext, documentID string) (rows int64, err error)
	DeleteLink(ctx RequestContext, id string) (rows int64, err error)
}

// ActivityStorer defines required methods for persisting document activity
type ActivityStorer interface {
	RecordUserActivity(ctx RequestContext, activity activity.UserActivity) (err error)
	GetDocumentActivity(ctx RequestContext, id string) (a []activity.DocumentActivity, err error)
	DeleteDocumentChangeActivity(ctx RequestContext, id string) (rows int64, err error)
}

// SearchStorer defines required methods for persisting search queries
type SearchStorer interface {
	IndexDocument(ctx RequestContext, doc doc.Document, a []attachment.Attachment) (err error)
	DeleteDocument(ctx RequestContext, ID string) (err error)
	IndexContent(ctx RequestContext, p page.Page) (err error)
	DeleteContent(ctx RequestContext, pageID string) (err error)
	Documents(ctx RequestContext, q search.QueryOptions) (results []search.QueryResult, err error)
}

// Indexer defines required methods for managing search indexing process
type Indexer interface {
	IndexDocument(ctx RequestContext, d doc.Document, a []attachment.Attachment)
	DeleteDocument(ctx RequestContext, ID string)
	IndexContent(ctx RequestContext, p page.Page)
	DeleteContent(ctx RequestContext, pageID string)
}

// BlockStorer defines required methods for persisting reusable content blocks
type BlockStorer interface {
	Add(ctx RequestContext, b block.Block) (err error)
	Get(ctx RequestContext, id string) (b block.Block, err error)
	GetBySpace(ctx RequestContext, spaceID string) (b []block.Block, err error)
	IncrementUsage(ctx RequestContext, id string) (err error)
	DecrementUsage(ctx RequestContext, id string) (err error)
	RemoveReference(ctx RequestContext, id string) (err error)
	Update(ctx RequestContext, b block.Block) (err error)
	Delete(ctx RequestContext, id string) (rows int64, err error)
}

// PageStorer defines required methods for persisting document pages
type PageStorer interface {
	Add(ctx RequestContext, model page.NewPage) (err error)
	Get(ctx RequestContext, pageID string) (p page.Page, err error)
	GetPages(ctx RequestContext, documentID string) (p []page.Page, err error)
	GetUnpublishedPages(ctx RequestContext, documentID string) (p []page.Page, err error)
	GetPagesWithoutContent(ctx RequestContext, documentID string) (pages []page.Page, err error)
	Update(ctx RequestContext, page page.Page, refID, userID string, skipRevision bool) (err error)
	Delete(ctx RequestContext, documentID, pageID string) (rows int64, err error)
	GetPageMeta(ctx RequestContext, pageID string) (meta page.Meta, err error)
	GetDocumentPageMeta(ctx RequestContext, documentID string, externalSourceOnly bool) (meta []page.Meta, err error)
	UpdateMeta(ctx RequestContext, meta page.Meta, updateUserID bool) (err error)
	UpdateSequence(ctx RequestContext, documentID, pageID string, sequence float64) (err error)
	UpdateLevel(ctx RequestContext, documentID, pageID string, level int) (err error)
	UpdateLevelSequence(ctx RequestContext, documentID, pageID string, level int, sequence float64) (err error)
	GetNextPageSequence(ctx RequestContext, documentID string) (maxSeq float64, err error)
	GetPageRevision(ctx RequestContext, revisionID string) (revision page.Revision, err error)
	GetPageRevisions(ctx RequestContext, pageID string) (revisions []page.Revision, err error)
	GetDocumentRevisions(ctx RequestContext, documentID string) (revisions []page.Revision, err error)
	DeletePageRevisions(ctx RequestContext, pageID string) (rows int64, err error)
}

// GroupStorer defines required methods for persisting user groups and memberships
type GroupStorer interface {
	Add(ctx RequestContext, g group.Group) (err error)
	Get(ctx RequestContext, refID string) (g group.Group, err error)
	GetAll(ctx RequestContext) (g []group.Group, err error)
	Update(ctx RequestContext, g group.Group) (err error)
	Delete(ctx RequestContext, refID string) (rows int64, err error)
	GetGroupMembers(ctx RequestContext, groupID string) (m []group.Member, err error)
	GetMembers(ctx RequestContext) (r []group.Record, err error)
	JoinGroup(ctx RequestContext, groupID, userID string) (err error)
	LeaveGroup(ctx RequestContext, groupID, userID string) (err error)
}
