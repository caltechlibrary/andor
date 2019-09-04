
## How this stuff works

And/Or JSON web service mapping URL end points to dataset[^2] commands.
It also includes a BasicAUTH authentication mechanism (for demo purposes),
and a simple user/role/object state permission scheme suitable for creating
role based workflows.

The And/Or service accepts HTTP GET and POST requests. If the request
is sent to a path with the collection name then And/Or will response
with an appropriate HTTP status and JSON payload or message. Before
returning the response a users' role permissions are applied. If a user
is unauthenticated then the BasicAUTH mechanism will trigger an
authentication request. If the And/Or prototype is successful BasicAUTH
will be replaced with something more appropriate like JSON Web Tokens
to control access.

And/Or doesn't provide user friendly user interface. It just speaks JSON.
A user interface is created with standard HTML, CSS and JavaScript. The
JavaScript is responsibly for making requests and retrieving appropriate
JSON objects and for formatting the results for us humans to read.

Currently the demo hosts one collection--- people. The people collection
is basic on Caltech Library People Identify project which is currently
curated via a Google Spreadsheet. This demo shows what curation might look
like in a more record oriented fashion.

The demo is based on the following workflow - depositors depositing
people records, a reviewer reviewing the deposits, a curator enhancing
records and determining if they should be "public". A publisher role
is also included to function like an application administrator (e.g.
"admin" in EPrints). 

Objects in our collection can be in one of the following states

deposit
: The initial state an object is created in before it is ready for review

review
: The review state is where curation happens, an object is edited as needed and sent to deleted, deposit, embargoed or public states

deleted
: Object is flagged for deletion from the collection

emabarged
: The object is ready for public view at sometime in the future

public
: The public should be able to read this object

The users can have one or more of the following roles--

publisher
: The collection administrator. They can create, read, update, delete objects, can change an object to any state.

curator
: Can read, update and delete objects. They can change an object's state to deposit, review, embargoed, published

depositor
: Can create/read/update/delete an object in deposit state, can send the object to review state.


A Role describes the permissions for objects and the states an 
object can be placed into.  Users can have more than one role.

This demo configuration creates the following users and roles

+ ester (publisher, public)
+ jane (depositor, curator, public)
+ millie (depositor, public)

The demo users have the same password, "hello". If you're 
running demo using the built-in BasicAUTH then use a private 
browser window to switch between logins.


[^1]: And/Or is named after a character in the [Ruby](https://www.zbs.org/index_new.php/store/ruby) stories produced by [ZBS](https://www.zbs.org) 

[^2]: [dataset](https://github.com/caltechlibrary/dataset) is a JSON object manager used from the command line or from Python scripts. It supports creating, reading, updating and delete objects in a collection. It has additional features like data frames.
