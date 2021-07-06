# TS-Tree
Implementation of a TS Tree for WORM data (i.e. for fact-totem).

Based on the white paper "TS-Trees: A Non-Alterable Search Tree Index for Trustworthy Databases on
Write-Once-Read-Many (WORM) Storage∗" by Jian Pei, Man Ki Mag Lau, Philip S. Yu.

Properties of a good index for Trustworthy Databases (WORM):

1. When committing a record to WORM storage, the id and path to access must be committed and immutable
2. Index must be incrementally maintainable, and scale to large databases
3. Data must be searchable using index efficiently
4. Cost of index space must be reasonable
5. Record must be destroyed once it expires

Goals of TS Trees:

1. Maintain virtual sorted list of records
2. Relatively balanced data structure
3. Provide effective mechanism to detect adversarial changes

Alternatives: General Hashting Trees (GHT), non-ordered and allows probing queries but no range queries

## Intent

TS-Tree tracks the keys and the _path_ to the value.  The Value blobs are stored in blob storage, and the semantics are
for Append only operations.  I.e. key/paths can be added but not altered or removed.