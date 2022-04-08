# auto-clicker

**auto-clicker** is an application that allows you to automate clicks with the left mouse button with the ability to adjust the coordinates of each click and the delay intervals between clicks.

## Dependencies

Application requires go version 1.17 to be installed

To work with the mouse, the [robotgo](https://github.com/go-vgo/robotgo) library is used.
Please see its requirements on the library page.

## Configuration

All available settings can be made using the configuration file `./config.json`.
An example configuration is in the file `./config.json.example`

To start using the application, you need to determine the coordinates where clicks will be made.
To determine the coordinates, you need to run the application with the `cfg` parameter, for example:

    go run main.go cfg

After that, move the cursor to the point on the screen where you will need to click in the future.
Every 5 seconds the application will read the current cursor coordinates and display them in the console.

> **Attention!** During the determination of the coordinates of real clicks will not be made.
> Also, when determining the coordinates, there is no need to make a click yourself.

To stop the process of reading coordinates, you need to transfer focus to the console and press the `q` key.
If you do not interrupt the process of determining the coordinates, it will stop on its own after 10 cycles.
If you need to define more coordinates, just run the command again.

After the command is processed, the `config.raw.json` file is generated in the root of the project
To use it in the application, you need to rename the file to `config.json` and make all the necessary changes to it.

> **Attention!** Running the configuration command again will overwrite the `config.raw.json` file!

### Configuration Options

    "name" - arbitrary name of the action. By default, the sequence number of the action is set, but can be changed to any name;
    "xCoordinate" - the x-coordinate for the click. Can be filled in automatically when configuring and reading coordinates;
    "yCoordinate" - y-coordinate for the click. Can be filled in automatically when configuring and reading coordinates;
    "delayAfter" - required delay in milliseconds before performing the next action. Filled in manually;
    "scrollAfterDelay" - the number of pixels to scroll after a click is made, and a delay if it was set.
        If you specify a negative value, it will scroll up. Filled in manually

## Start

Once the `./config.json` file has been configured, the application can be run.
To do this, you need to start the application with the `run` parameter, for example:

    go run main.go run

To stop the program execution process, you need to transfer focus to the console and press the `q` key
If you do not interrupt the process of determining the coordinates, it will stop on its own after 30 000 cycles.

> **Attention!**
> If the configuration file does not specify values for the delay: `delayAfter` -
> it can be very problematic to stop the program execution process, 
> because the mouse cursor will move very quickly along the coordinates specified in the configuration file.
