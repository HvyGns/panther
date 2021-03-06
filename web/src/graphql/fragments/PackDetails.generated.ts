/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import * as Types from '../../../__generated__/schema';

import { GraphQLError } from 'graphql';
import gql from 'graphql-tag';

export type PackDetails = Pick<
  Types.Pack,
  | 'id'
  | 'displayName'
  | 'description'
  | 'enabled'
  | 'updateAvailable'
  | 'lastModified'
  | 'lastModifiedBy'
  | 'createdAt'
  | 'createdBy'
> & {
  packVersion: Pick<Types.PackVersion, 'id' | 'name'>;
  availableVersions: Array<Pick<Types.PackVersion, 'id' | 'name'>>;
  detectionTypes: Pick<Types.DetectionTypes, 'GLOBAL' | 'POLICY' | 'RULE'>;
  detectionsPatterns: Pick<Types.PackDetectionsPatterns, 'IDs'>;
};

export const PackDetails = gql`
  fragment PackDetails on Pack {
    id
    displayName
    description
    enabled
    updateAvailable
    packVersion {
      id
      name
    }
    availableVersions {
      id
      name
    }
    detectionTypes {
      GLOBAL
      POLICY
      RULE
    }
    detectionsPatterns {
      IDs
    }
    lastModified
    lastModifiedBy
    createdAt
    createdBy
  }
`;
