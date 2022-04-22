# Keys

Macrometa implements the ArangoDB interface and, with ith, several constrains on keys apply. 
It is important to realize that "key" also applies to collection names, document names, graph,
and edge names. Therefore, the rules below apply whenever an identifier needs to be defined:

Users can pick key values as required, provided that the values conform to the following restrictions:

* The key must be a string value. Numeric keys are not allowed, but any numeric value can be put into a string and can then be used as document key.
* The key must be at least 1 byte and at most 254 bytes long. Empty keys are disallowed when specified (though it may be valid to completely omit the _key attribute from a document)
*  It must consist of the letters a-z (lower or upper case), the digits 0-9 or any of the following punctuation characters: _ - : . @ ( ) + , = ; $ ! * ' %
*  Any other characters, especially multi-byte UTF-8 sequences, whitespace or punctuation characters cannot be used inside key values
* The key must be unique within the collection it is used
* Keys are case-sensitive, i.e. myKey and MyKEY are considered to be different keys.

Specifying a document key is optional when creating new documents. 
If no document key is specified by the user, ArangoDB will create the document key itself as each document is required to have a key.
