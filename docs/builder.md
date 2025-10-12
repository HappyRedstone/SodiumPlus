# Builder Tool

This project includes a command-line tool for building and managing the modpack.

## Building the tool

To build the tool, run the following command:

```bash
make bootstrap
```

This will create binaries for different operating systems and architectures in the `bin` directory.

## Running the tool

You can run the tool from the `bin` directory. For example, on a 64-bit Linux system, you would use:

```bash
./bin/builder.x86_64 <command>
```

## Commands

Here are some of the available commands:

*   `list` (aliases: `l`, `ls`): Lists all the mods in the modpack.
    *   `--json`, `-j`: Output as JSON.
    *   `--html`, `-H`: Output as HTML.
    *   `--markdown`, `-m`: Output as Markdown.
    *   `--console`, `-c`: Output in a console-friendly format (default).
    *   `--output`, `-o`: Specify an output file.
*   `client`: Contains subcommands for managing the client-side pack.
*   `server`: Contains subcommands for managing the server-side pack.
*   `version`: Contains subcommands for managing different versions of the modpack.
*   `serve`: Serves a local website for viewing pack information.
*   `install`: Installs the modpack.
*   `clean-all`: Cleans all generated files.
*   `bundle-all`: Bundles all versions of the modpack.

## Example: Listing Mods for a Specific Version

To list all the mods for a specific version of the modpack (e.g., `1.21.10` with the `Fabric` loader), you first need to bundle that version, and then you can list the mods from the generated pack file.

### 1. Bundle the version

This command will create the necessary pack files in the `.build/temp` directory.

```bash
./bin/builder.x86_64 version bundle 1.21.10 Fabric
```

### 2. List the mods

Once the version is bundled, you can use the `list` command with the path to the generated `pack.toml` file.

```bash
./bin/builder.x86_64 --pack-file .build/temp/1.21.10/Fabric/pack.toml list
```
