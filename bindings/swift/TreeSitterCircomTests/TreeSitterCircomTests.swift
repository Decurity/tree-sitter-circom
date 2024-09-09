import XCTest
import SwiftTreeSitter
import TreeSitterCircom

final class TreeSitterCircomTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_circom())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Circom grammar")
    }
}
