# -------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
# --------------------------------------------------------------------------

from cmath import exp
from apistub.nodes import TypeHintParser


class TestTypeHintParser:

    def test_typehint(self):
        parser = TypeHintParser([])
        code = """
        # type: (str) -> str
        return val
        """
        expected = "str"
        assert parser._parse_typehint(code) == expected

    def test_typehint_no_spaces(self):
        parser = TypeHintParser([])
        code = """
        # type:(str)->str
        return val
        """
        expected = "str"
        assert parser._parse_typehint(code) == expected

    def test_typehint_with_pylint_disable(self):
        parser = TypeHintParser([])
        code = """
        # type: (...) -> AnalyzeHealthcareEntitiesLROPoller[ItemPaged[Union[AnalyzeHealthcareEntitiesResult, DocumentError]]]  # pylint: disable=line-too-long
        """
        expected = "AnalyzeHealthcareEntitiesLROPoller[ItemPaged[Union[AnalyzeHealthcareEntitiesResult, DocumentError]]]"
        assert parser._parse_typehint(code) == expected
