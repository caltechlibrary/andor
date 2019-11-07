/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */




/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


// error_clear will set the global error state to nil.
//

extern void error_clear();

// error_message returns an error message previously recorded or
// an empty string if no errors recorded
//

extern char* error_message();

// use_strict_dotpath sets the library option value for
// enforcing strict dotpaths. 1 is true, any other value is false.
//

extern int use_strict_dotpath(int p0);

// is_verbose returns the library options' verbose value.
//

extern int is_verbose();

// verbose_on set library verbose to true
//

extern void verbose_on();

// verbose_off set library verbose to false
//

extern void verbose_off();

// dataset_version returns the dataset version the libdataset presents.
//

extern char* dataset_version();

// init_collection intializes a collection and records as much metadata
// as it can from the execution environment (e.g. username,
// datetime created)
//

extern int init_collection(char* p0);

// create_object takes JSON source and adds it to the collection with
// the provided key.
//

extern int create_object(char* p0, char* p1, char* p2);

// read_object takes a key and returns JSON source of the record
//

extern char* read_object(char* p0, char* p1, int p2);

// THIS IS AN UGLY HACK, Python ctypes doesn't **easily** support
// undemensioned arrays of strings. So we will assume the array of
// keys has already been transformed into JSON before calling
// read_list.
//

extern char* read_object_list(char* p0, char* p1, int p2);

// update_object takes a key and JSON source and replaces the record
// in the collection.
//

extern int update_object(char* p0, char* p1, char* p2);

// delete_object takes a key and removes a record from the collection
//

extern int delete_object(char* p0, char* p1);

// join takes a collection name, a key, and merges JSON source with an
// existing JSON record. If overwrite is 1 it overwrites and replaces
// common values, if not 1 it only adds missing attributes.
//

extern int join(char* p0, char* p1, char* p2, int p3);

// key_exists returns 1 if the key exists in a collection or 0 if not.
//

extern int key_exists(char* p0, char* p1);

// keys returns JSON source of an array of keys from the collection
//

extern char* keys(char* p0);

// key_filter returns JSON source of an array of keys passing
// through the filter of objects in the collection.
//

extern char* key_filter(char* p0, char* p1, char* p2);

// key_sort returns JSON source of an array of keys sorted by
// the sort expression applied to the objects in the collection.
//

extern char* key_sort(char* p0, char* p1, char* p2);

// count returns the number of objects (records) in a collection.
// if an error is encounter a -1 is returned.

extern int count(char* p0);

// import_csv - import a CSV file into a collection
// syntax: COLLECTION CSV_FILENAME ID_COL
//
// options that should support sensible defaults:
//
//     cUseHeaderRow
//     cOverwrite
//

extern int import_csv(char* p0, char* p1, int p2, int p3, int p4);

// export_csv - export collection objects to a CSV file
// syntax: COLLECTION FRAME CSV_FILENAME
//

extern int export_csv(char* p0, char* p1, char* p2);

// import_gsheet - import a GSheet into a collection
// syntax: COLLECTION GSHEET_ID SHEET_NAME ID_COL CELL_RANGE
//
// options that should support sensible defaults:
//
//    cUseHeaderRow
//    cOverwrite
//

extern int import_gsheet(char* p0, char* p1, char* p2, int p3, char* p4, int p5, int p6);

// export_gsheet - export collection objects to a GSheet
// syntax: COLLECTION FRAME GSHEET_ID GSHEET_NAME CELL_RANGE
//

extern int export_gsheet(char* p0, char* p1, char* p2, char* p3, char* p4);

// status checks to see if a collection exists or not.
//

extern int status(char* p0);

// list returns JSON array of objects in a collections based on a
// JSON array of keys.
//

extern char* list(char* p0, char* p1);

// path returns the path on disc to an JSON object document
// in the collection.
//

extern char* path(char* p0, char* p1);

// check runs the analyzer over a collection and looks for
// problem records.
//

extern int check(char* p0);

// repair runs the analyzer over a collection and repairs JSON
// objects and attachment discovered having a problem. Also is
// useful for upgrading a collection between dataset releases.
//

extern int repair(char* p0);

// attach will attach a file to a JSON object in a collection. It takes
// a semver string (e.g. v0.0.1) and associates that with where it stores
// the file.  If semver is v0.0.0 it is considered unversioned, if v0.0.1
// or larger it is considered versioned.
//

extern int attach(char* p0, char* p1, char* p2, char* p3);

// attachments returns a list of attachments and their size in
// associated with a JSON obejct in the collection.
//

extern char* attachments(char* p0, char* p1);

// detach exports the file associated with the semver from the JSON
// object in the collection. The file remains "attached".
//

extern int detach(char* p0, char* p1, char* p2, char* p3);

// prune removes an attachment by semver from a JSON object in the
// collection. This is destructive, the file is removed from disc.
//

extern int prune(char* p0, char* p1, char* p2, char* p3);

// clone takes a collection name, a JSON array of keys and creates
// a new collection with a new name based on the origin's collections'
// objects.
//

extern int clone(char* p0, char* p1, char* p2);

// clone_sample is like clone both generates a sample or test and
// training set of sampled of the cloned collection.
//

extern int clone_sample(char* p0, char* p1, char* p2, int p3);

// grid generates a "Grid" structure from a collection.
//

extern char* grid(char* p0, char* p1, char* p2);

// frame_exists returns 1 (true) if frame name exists in collection, 0 (false) otherwise
//

extern int frame_exists(char* p0, char* p1);

// frame_keys takes a collection name and frame name and returns a list of keys from the frame or an empty list.
// The list is expressed as a JSON source.
//

extern char* frame_keys(char* p0, char* p1);

// frame_create defines a new frame an populates it.
//

extern int frame_create(char* p0, char* p1, char* p2, char* p3, char* p4);

// frame_objects retrieves a JSON source list of objects from a frame.
//

extern char* frame_objects(char* p0, char* p1);

// frame_refresh refresh the contents of a frame given a list of keys.
//

extern int frame_refresh(char* p0, char* p1, char* p2);

// frame_reframe will change of object list in a frame based on the key list provided.
//

extern int frame_reframe(char* p0, char* p1, char* p2);

// frame_clear will clear the object list and keys associated with a frame.
//

extern int frame_clear(char* p0, char* p1);

// frame_delete will removes a frame from a collection
//

extern int frame_delete(char* p0, char* p1);

// frames returns a JSON array of frames names in the collection.
//

extern char* frames(char* p0);

// sync_send_csv - synchronize a frame sending data to a CSV file
//

extern int sync_send_csv(char* p0, char* p1, char* p2, int p3);

// sync_recieve_csv - synchronize a frame recieving data from a CSV file
//

extern int sync_recieve_csv(char* p0, char* p1, char* p2, int p3);

// sync_send - synchronize a frame sending data to a GSheet
//

extern int sync_send_gsheet(char* p0, char* p1, char* p2, char* p3, char* p4, int p5);

// sync_recieve_gsheet - synchronize a frame recieving data from a GSheet
//

extern int sync_recieve_gsheet(char* p0, char* p1, char* p2, char* p3, char* p4, int p5);

// frame_grid takes a frames object list and returns a grid
// (2D JSON array) representation of the object list.
// If the "header row" value is 1 a header row of labels is
// included, otherwise it is only the values of returned in the grid.
//

extern char* frame_grid(char* p0, char* p1, int p2);

//
// make_objects - is a function to creates empty a objects in batch.
// It requires a JSON list of keys to create. For each key present
// an attempt is made to create a new empty object based on the JSON
// provided (e.g. `{}`, `{"is_empty": true}`). The reason to do this
// is that it means the collection.json file is updated once for the
// whole call and that the keys are now reserved to be updated separately.
// Returns 1 on success, 0 if errors encountered.
//

extern int make_objects(char* p0, char* p1, char* p2);

//
// update_objects - is a function to update objects in batch.
// It requires a JSON array of keys and a JSON array of
// matching objects. The list of keys and objects are processed
// together with calls to update individual records. Returns 1 on
// success, 0 on error.
//

extern int update_objects(char* p0, char* p1, char* p2);

// set_who will set the "who" value associated with the collection's metadata
//

extern int set_who(char* p0, char* p1);

// get_who will get the "who" value associated with the collection's metadata
//

extern char* get_who(char* p0);

// set_what will set the "what" value associated with the collection's metadata
//

extern int set_what(char* p0, char* p1);

// get_what will get the "what" value associated with the collection's metadata
//

extern char* get_what(char* p0);

// set_when will set the "when" value associated with the collection's metadata
//

extern int set_when(char* p0, char* p1);

// get_when will get the "what" value associated with the collection's metadata
//

extern char* get_when(char* p0);

// set_where will set the "where" value associated with the collection's metadata
//

extern int set_where(char* p0, char* p1);

// get_where will get the "where" value associated with the collection's metadata
//

extern char* get_where(char* p0);

// set_version will set the "version" value associated with the collection's metadata
//

extern int set_version(char* p0, char* p1);

// get_version will get the "version" value associated with the collection's metadata
//

extern char* get_version(char* p0);

// set_contact will set the "contact" value associated with the collection's metadata
//

extern int set_contact(char* p0, char* p1);

// get_contact will get the "contact" value associated with the collection's metadata
//

extern char* get_contact(char* p0);

#ifdef __cplusplus
}
#endif
