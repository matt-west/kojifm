# Koji FM #
A web service that generates a playlist of youtube videos based on your last.fm history.

This project was previously written in Ruby on Rails however I am re-writing it in GO so that I can get to know the language.

### How to get going ###
1. Compile the source - 8g kojifm.go
2. Link it - 8l kojifm.8
3. Run it - ./8.out

These commands are for a linux system you may need to use 6g & 6l instead. You should know which command to use from the installation of GO.

Currently the code uses http.ListenAndServe to start up the server, this will be changed once everything is ready for AppEngine.

### ToDo/Wishlist ###
- Add authentication so that views can be scrobbled back to last.fm
- Make it so that playlists can be saved to a user's youtube account
