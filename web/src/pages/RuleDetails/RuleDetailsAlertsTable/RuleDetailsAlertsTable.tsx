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

import React from 'react';
import { AbstractButton, Box, Link, Table } from 'pouncejs';
import { Link as RRLink } from 'react-router-dom';
import urls from 'Source/urls';
import { formatDatetime, shortenId } from 'Helpers/utils';
import { ListAlertsForRule } from '../graphql/listAlertsForRule.generated';
import SeverityBadge from 'Components/SeverityBadge';

interface RuleDetailsAlertsTableProps {
  alerts: ListAlertsForRule['alerts']['alertSummaries'];
}

const RuleDetailsAlertsTable: React.FC<RuleDetailsAlertsTableProps> = ({ alerts }) => {
  return (
    <Table>
      <Table.Head>
        <Table.Row>
          <Table.HeaderCell>Title</Table.HeaderCell>
          <Table.HeaderCell>Created At</Table.HeaderCell>
          <Table.HeaderCell>Alert ID</Table.HeaderCell>
          <Table.HeaderCell align="right">Events</Table.HeaderCell>
          <Table.HeaderCell>Last Matched At</Table.HeaderCell>
        </Table.Row>
      </Table.Head>
      <Table.Body>
        {alerts.map(alert => (
          <Table.Row key={alert.alertId}>
            <Table.Cell maxWidth={400} truncated title={alert.title}>
              <Link as={RRLink} to={urls.logAnalysis.alerts.details(alert.alertId)} py={4} pr={4}>
                {alert.title} #{shortenId(alert.alertId)}
              </Link>
            </Table.Cell>
            <Table.Cell align="right">{alert.eventsMatched}</Table.Cell>
            <Table.Cell align="center">
              <Box my={-1} display="inline-block">
                <SeverityBadge severity={alert.severity} />
              </Box>
            </Table.Cell>
            <Table.Cell align="right">{formatDatetime(alert.creationTime)}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table>
  );
};

export default RuleDetailsAlertsTable;
