# Aya - Simple Version Control System

Aya is a simple and easy-to-use version control system that focuses on simplicity and usability. Unlike Git, which has a myriad of commands and complex error messages, Aya offers a minimal set of commands, ensuring a fast learning curve and preventing confusing errors. Our goal is to make version control accessible and straightforward for everyone.

## Key Features

- Less than 6 commands
- Intuitive and user-friendly
- Focuses on simplicity over functionality
- Fast learning curve
- Prevents complex errors like those in Git

## Example Error in Git
Here's an example of a confusing error in Git that you won't encounter in Aya:
```sh
! [rejected] main -> main (fetch first)
error: failed to push some refs to 'https://github.com/Mohamedjcali/aya'
hint: Updates were rejected because the remote contains work that you do not
hint: have locally. This is usually caused by another repository pushing to
hint: the same ref. If you want to integrate the remote changes, use
hint: 'git pull' before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.
```

## Installation

### Download the Code

1. Go to the [code](https://github.com/Mohamedjcali/aya) page.
2. Download the code.
3. Build the code using `go build`.
4. You will see an `aya.exe` file in that directory.

### Windows

#### Add Directory to PATH:

1. Right-click on This PC or Computer on the desktop or in File Explorer.
2. Click on Properties.
3. Click on Advanced system settings.
4. Click on the Environment Variables button.
5. In the System variables section, find the Path variable, select it, and click Edit.
6. Click New and add the directory where you placed your .exe file (e.g., `C:\Tools`).
7. Click OK to close all dialog boxes.

### Linux

1. Ensure your project is built and you have the executable. If not, use the following command to build it:
    ```sh
    go build .
    ```
2. Move the executable to a directory that's in your PATH. Common directories include `/usr/local/bin` or `/usr/bin`.
3. If you choose to create your own directory, you need to add it to the PATH:
    - Open a terminal.
    - Edit your shell profile file (e.g., `.bashrc`, `.bash_profile`, `.zshrc`, etc.):
        ```sh
        nano ~/.bashrc  # or whichever file your shell uses
        ```
    - Add the following line at the end of the file:
        ```sh
        export PATH=$PATH:/path/to/your/directory
        ```
    - Save the file and reload the profile:
        ```sh
        source ~/.bashrc  # or whichever file your shell uses
        ```
4. Test:
    - Open a new terminal window and type `aya --help` to ensure it works from any location:
        ```sh
        aya -h
        ```

By following these steps, you'll be able to run Aya commands from anywhere on your computer, both on Windows and Linux.

## Usage

### Checking Installation

To check if Aya is installed correctly, open your terminal and run:

```sh
aya -h
```
If you see an error message like:


```sh
'aya' is not recognized as an internal or external command, operable program or batch file.
```
then Aya is not installed correctly. Please follow the installation instructions again.

If Aya is installed correctly, you will see a help message explaining the available commands.

Initializing a Project
To start using Aya, you need to initialize a project. Run the following command:

```sh
aya init
```
Aya will prompt you to enter the project name and basic information about the project:

```sh
Please enter the project name:
```
Type the project name and press Enter.

```sh
Please enter the project basic info:
```
Type a short description of the project and press Enter.

Aya will create a hidden .aya folder. To view hidden folders, go to your file explorer, look at the taskbar, go to the View tab, and enable the option to show hidden folders. Aya uses this folder to store all project-related information.

## Adding a New Version
To add a new version to your project, use the aya add command with the necessary flags. If you run aya add without any flags, you will get an error because certain flags are required.

Here are the flags you need to use:

-v : The version name. It must contain only numbers and dots (e.g., 2.0.0.1). If you enter an invalid version name (e.g., 2.0.0d), Aya will give you an error.
-w : The writer's name. This flag is required.
-d : The description of the version. Enclose the description in quotes.
-p : The branch to save the version in. This flag is optional. If not provided, Aya will use the branch specified in the .aya-config file.
Example usage:

```sh
aya add -v 0.0.0.1 -w cqani -d "This is the description of the version" -p main
```
## Loading a Version
To load a version, use the aya load command. By default, Aya will load the latest saved version:

```sh
aya load
```
This command loads the latest version from the .aya/refs folder.

To load a specific version, use the -v flag:

```sh
aya load -v 1.0.0
```
To specify a branch, use the -p flag:

```sh
aya load -v 1.0.0 -p main
```
## Copying a Repository
The aya copy command works like git clone. It copies a given GitHub repository to a folder created with the same name as the repository.

Example usage:

```sh
aya copy https://github.com/user/repository.git
```
To specify a destination folder:
```sh
aya copy https://github.com/user/repository.git destination
```
Conclusion
Aya is designed to be a simple and easy-to-use version control system. By following the instructions above, you should be able to get started with Aya and manage your projects efficiently.
also still aya push and pull they are in work when we finish it we going to puplish it
