## Javascript File Upload Demo

A small demo that utilizes HTML5 API's and Javascript to provide a custom file upload UI.

The demo follows along with the following MDN article.

https://developer.mozilla.org/en-US/docs/Using_files_from_web_applications

It includes code to:
 - examine selected files client side (name, size and type) plus show thumbnail of images
 - use a custom link (instead of standard form button) to trigger file upload window
 - sets up a basic drag and drop zone for selecting files
 - upload selected files asynchronously, including feedback display through the UI
 - includes server side code using Python & Flask

## Usage

### Python Server App

1. Install Python 3.5
2. Run `pyvenv env`.
3. Run `. ./env/bin/activate`.
4. Run `pip install -r requirements.txt`.
5. Run `python server.py`.

The URL to access in your browser will be printed out, but by default the app listens on port 5001 so http://localhost:5001/ should work.

### Javascript

You can also just open `index.html` in your browser directly.  This will work for all of the functionality except the file upload part.  You need the server running to test that.
