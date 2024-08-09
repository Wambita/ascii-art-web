# ASCII-ART WEB

ASCII Art Web is a dynamic web application that allows users to create ASCII ART.It consists of creating and running a server by which a web-based GUI (graphical user interface) implementation of the previous project ascii-art.


# TABLE OF CONTENT

- [Introduction](#Introduction)
- [Usage](#Usage)
- [Implementation](#implementation-detail:algorithm)
- [Authors](#author)


## Introduction
Ascii-art Web is a project that hosts webpages for users to generate Ascii-Art of their desired text with their desired banner. The user is presented an area for them to type their desired text. If the character is among the printable ASCII character range  it is displayed this includes,numerics and symbols.

### Features 
+ Convert text to ASCII art using varius banner files
+ Interactive web interface for  creating ASCII art.
+ Responsive design for desktop and mobile devices.


## Installation 

Download the Go installer from the official [Go website](https://go.dev/doc/install).

Follow the installation instructions for your specific operating system.

Verify your installation by opening a command prompt and typing `go --version`. This should print the installed version of Go.


To set up ASCII Art Web locally, follow these steps:
```bash
# Clone the repository
git clone https://learn.zone01kisumu.ke/git/shfana/ascii-art-web.git

# Navigate to the project directory
cd ascii-art-web

# Install dependencies

`sudo apt install go -all`

# Start the ascii-web server
go run .
```

## Usage
After installation, you can access the web application by navgating to [http://localhost:8080](http://localhost:8080)



You will receive the following message on your terminal
```bash
Server is starting at http://localhost:8080
```
![home](/assets/images/home.png)
 

- These are the various parts of the home page 

![instructions](/assets/images/instruct.png)


 + Navbar - The navbar links to other pages of the application

 + The Input section - where the user keys in the input they want to generate to art

 + banner selection - The user has three options of banner files to choose from. (shadow, thinkertoy and standard)

 + The submit button - The user presses this button to generate the art that is displayed in the output section.

 + The footer  - This section contains the copyright disclaimer and an icon that has a  link to the repository.


![banner](/assets/images/banner.png)


The user can select one of the three available fonts/banner files to display their input in art form.

This project comes with three predefined ASCII fonts, located in banner files:

  - [`shadow`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
  - [`standard`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
  - [`thinkertoy`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)


**- Each character is represented over 8 lines, providing a clear and sizable output.**

**- Characters are separated by a new line `\n`.**

### Creating ASCII Art

   1. Enter your desired text in the input field
   2. Select a banner style from the dropdown menu
   3. Click "Generate" to create your ASCII art


- From the example where the user has typed "hello" in the input field and selected  the `standard` banner. 
- The output should be shown as seen in the image below
![output](/assets/images/outputS.png)

- In the case of input `Hello` and `shadow` banner selected, the output is as shown below
![shadow](/assets/images/ouputSh.png)


- In the case of input `Hello` and `thinkertoy` banner selected.
![thinkertoyinput](/assets/images/thinkertoy%20input.png)

- Output is as shown below
![thnkertoyoutput](/assets/images/thinkertou%20output.png)

- Bad Requests: They occur when

    - The user does not give an input
    - The user gives an input that is not within the printable ASCII character range (from space to tide `~`). There are 95 characters in total

An example of a bad request is shown below(the user tries to input chinese characters or emojis)
![badrequest](/assets/images/badrequest.png)

The Following error is returned.
![badreq](/assets/images/badrequestout.png)

The user can click `go back` to return home to try using the program again


- In the case where a user tries to navigate to a non existent page or use a banner file that is not  available, he/she gets a 404 not found error 
![notfound](/assets/images/not%20found.png)

- Internal server error occurs when an error occurs from the program and is not a user error.
![internalservererror](/assets/images/ISE.png)

## Implementation detail: algorithm

Once you run the program in the using the command `go run .` the program runs on a localhost server with the port provided. When you open the localhost provided you are directed to a page which resembles the image below. 

![alt text](assets/images/input.png)

One the text box you enter the words you would like to be turned into word art. Then below it there is a dropdown option which helps you choose which type of bannerfile you can use. We have three options: shadow, thinkertoy and standard.
Once you have selected the banner file you can click on the submit button to get your word art. This issues a get request to the server which intern takes the word type in the text area and the bannerfile selected as the inputs. When the program is run the following variations of processes occur:

### Case 1: Print non-printable character.
If the character doesn't range between the ascii values 32-126 and is not '\n' or '\t' then it is considered as a bad error and an error message is printed out.

![alt text](assets/images/badrequestout.png)

### Case 2: Trying to run invalid page.
If the user tries to run a page that is not available then a 404 not found error message is printed

![alt text](<assets/images/not found.png>)

### Case 3: Any other error
If the program encounters any error that is not specified, it returns an error message 500 internal server error.

![alt text](assets/images/ISE.png)

### Case 4: printable charachter.
If the character is printable then it will issue a POST request t the server to return the result of the program to the output textbox.

![alt text](assets/images/ouputSh.png)

When running the program one of the four will always occur deppending on what you have issued to the program.

## Contribution

Contributions are welcome. Please adhere to the existing coding standards and include unit tests for any new features or changes. Ensure to thoroughly test the code before pushing any updates.
If you encounter any issues or have suggestions for improvement, feel free to submit an issue, pull request or propose a change!


Please follow these steps to contribute:

   1. Fork the repository
   2. Create a new branch (git checkout -b feature/AmazingFeature)
   3. Commit your changes (git commit -m 'Add some AmazingFeature')
   4. Push to the branch (git push origin feature/AmazingFeature)
   5. Open a Pull Request


## License
This Program is under the **MIT** Licence see [LICENSE](/LICENSE) for details.

## Authors
This program was built and maintained by 
- [cliffootieno](https://learn.zone01kisumu.ke/git/cliffootieno)
- [wnjuguna](https://learn.zone01kisumu.ke/git/wnjuguna)
- [shfana](https://learn.zone01kisumu.ke/git/shfana)