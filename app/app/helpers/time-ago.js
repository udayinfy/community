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

import Ember from 'ember';
import dateUtil from '../utils/date';

// {{time-ago createdAt}}
export function timeAgo(params) {
    let d = dateUtil.timeAgo(params[0]);

	if (d === 'Invalid date'){
		d = '';
	}

	return d;
}

export default Ember.Helper.helper(timeAgo);
