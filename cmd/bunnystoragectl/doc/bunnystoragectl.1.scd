bunnystoragectl(1)

# NAME

bunnystoragectl - A command-line interface for Bunny Storage

# SYNOPSIS

*bunnystoragectl* [options...] [commands...] [options...]

# DESCRIPTION

*bunnystoragectl* is an easy-to-use command-line interface for Bunny Storage,
the edge storage service provided by Bunny.net.

# OPTIONS

*-z*, *--storage-zone*
	Provide the name of the storage zone in Bunny.net. Can also be set using the
	BUNNY_STORAGE_ZONE environment variable.

*-k*, *--key*
	Provide the read/write storage zone's API key. Can also be set using the
	BUNNY_KEY environment variable.

*-e*, *--endpoint*
	Provide the endpoint for the storage zone in Bunny.net. By default,
	"https://storage.bunnycdn.com" will be used. Can also be set using the
	BUNNY_ENDPOINT environment variable.

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version number and quit.

*-V*, *--verbose*
	Enable verbose output.

# COMMANDS

*ls*, *list* <options>
	List files and directories in the storage zone.

	Options are:

	*-p*, *--path*
		Set the directory path that you want to list. Defaults to the root
		directory.

	*-j*, *--json*
		Enable JSON output. Defaults to false.

*cp*, *upload* <options>
	Upload a file to a storage zone.

	Options are:

	*-p*, *--path*
        Set the directory path to where your file will be stored. Defaults to
        the root directory.

	*-f*, *--file*
		Set the path to the file you want to upload.

*dl*, *download* <options>
	Download a file from a storage zone.

	Options are:

	*-p*, *--path*
        Set the directory path to your file in Bunny.net. Defaults to the root
        directory.

	*-f*, *--file*
		Set the file you want to download.

	*-o*, *--output*
		Set the path and filename for the downloaded file.

	*-F*, *--force*
		Set whether existing files should be overwritten.

*rm*, *remove* <options>
	Remove a file from a storage zone.

	Options are:

	*-p*, *--path*
        Set the directory path to your file in Bunny.net. Defaults to the root
        directory.

	*-f*, *--file*
		Set the file you want to remove.

# EXAMPLES

*Example 1. List files*
	The following command line lists all files inside the "path/to/files/ path
	inside the storage zone "my-zone". It uses environment variables to provide
	*bunnystoragectl* with the credentials and information it needs.

	BUNNY_KEY='my-key' BUNNY_STORAGE_ZONE='my-zone' bunnystoragectl list -p 'path/to/files'

*Example 2. Download file*
	The following command line download a file called "filename" from the
	"path/to/files" path inside the storage zone "my-zone". It uses flags to
	provide *bunnystoragectl* with the credentials and information it needs.

	It's recommended to provide the storage zone key using environment variables
	instead, for better security.

	bunnystoragectl -k 'my-key' -z 'my-zone' download -F -p 'path/to/files' -f 'filename'

*Example 3. Upload file*
    The following command line upload a file called "file.txt" located at
    "/path/to/local/" in the local filesystem to the storage zone "my-zone" and
    path "path/to/files/in/bunny".

	As with the previous example, it uses flags to provide *bunnystoragectl*
	with the information it needs, but this time it provides the storage zone
	key using an environment variable for better security.

	BUNNY_KEY='my-key' bunnystoragectl -z 'my-zone' cp -p 'path/to/files/in/bunny' -f '/path/to/local/file.txt'

# ENVIRONMENT

$BUNNY_STORAGE_ZONE
	The storage zone from Bunny.net to use.

$BUNNY_KEY
	The storage zone API key to use.

$BUNNY_ENDPOINT
	The API endpoint from Bunny.net to use.

# REPORTING BUGS

Report bugs via email to <~jamesponddotco/bunnystorage@todo.sr.ht> or via the
web interface at https://todo.sr.ht/~jamesponddotco/bunnystorage.

# AUTHORS

Maintained by James Pond <james@cipher.host>.
