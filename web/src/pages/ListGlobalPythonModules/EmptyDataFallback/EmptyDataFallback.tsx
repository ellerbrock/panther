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
import { Box, Button, Flex, Heading, Text } from 'pouncejs';
import EmptyDataImg from 'Assets/illustrations/empty-box.svg';
import { Link as RRLink } from 'react-router-dom';
import urls from 'Source/urls';

const EmptyDataFallback: React.FC = () => (
  <Flex height="100%" width="100%" justify="center" align="center" direction="column">
    <Box m={10}>
      <img alt="Empty data illustration" src={EmptyDataImg} width="auto" height={400} />
    </Box>
    <Heading mb={6}>It{"'"}s empty in here</Heading>
    <Text color="gray-200" textAlign="center" mb={8}>
      Python modules allow you to define re-usable code elements for your rules and policies
    </Text>
    <Button as={RRLink} to={urls.settings.globalPythonModules.create()}>
      Create a Global Python Module
    </Button>
  </Flex>
);

export default EmptyDataFallback;
