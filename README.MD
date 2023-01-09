# Welcome to ASCII ART WEB PROJECT


# Project Description

Ascii-art-web consists of creating and running a server, in which it will be possible to use a web  **GUI**  (graphical user interface) version of the [ascii-art](https://01.gritlab.ax/git/root/public/src/branch/master/subjects/ascii-art).

The webpage must allow the use of the different banners:

-   [shadow](https://01.gritlab.ax/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
-   [standard](https://01.gritlab.ax/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
-   [thinkertoy](https://01.gritlab.ax/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)

This project was built with the following tools:

Golang
HTML
CSS


## Usage
To get a local copy up and running follow these simple example steps.

### Prerequisites
The following prerequisites are needed for running the program.

1.	Go Programming Language (golang). This can be downloaded from https://go.dev/dl/
2.	Golang can be installed by following the instructions from https://www.youtube.com/watch?v=nWZRZUg2lOM

### Installation
To run this project you will need to clone the following repository on your computer locally:

`git clone https://01.gritlab.ax/git/oluwatosin/ascii-art-web.git`

Change to the ascii-art-web directory, and run the project there:

    cd ascii-art-web

Open the terminal and run:

    go run main.go
Open a web browser, e.g. google chrome and in the address bar type:

> `localhost:8080/home`

press the enter key.
The page should load with the title ASCII ART GENERATOR

Enter any Ascii character in the input area and see the ascii art generated.

## Implementation Algorithm
1.  Input is received in form of strings and validated to ensure that they are all ascii characters. If non-ascii characters has been entered a message is printing informing user that only ascii characters are allowed.
2. Depending on the banner the user has chosen, the banner file is read into a slice of string.
3. The slice is splited using the newline, line feed character `(\r\n)`. The resulting slice is then looped through height of 8 and the content of the banner file is written to string.

## Authors

Mathisen (cat lover)
Oluwatosin (fish lover)

