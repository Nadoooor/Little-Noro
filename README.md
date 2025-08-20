# Little-Noro
A GUI Pet like a Tamagotchi that can be treated and eats unwanted files.

## About 
It is a small pet where you can feed him and level him to open up new pets. It has sound effects when you feed it or when level up. 
Also, It has an easter egg go ahead and look for the easter egg. And time based wallpaper.(Hint: Don't cheat)<br>

> [!NOTE]
>You have 4 Levels and it debends on your points. (points > 20 : level 1, points <= 20 and > 40 : level 2, and so on)
## techniques used to Optimize.
- Loaded all the images first in a fyne.Resource so it doesn't talk alot of ram when called
- Checking every second if the points or pet or clock, etc... changed and it won't redraw the widgets unless they change.
- Made save and load functions to save to or load from the JSON file and reusing them instead of reading the json every time.
- Made the wallpaper size to be 600X400 by defualt instead of resizeing it with fyne so it won't use a lot of memory.


## how to install ??
Download the ZIP file from the release and extract it.<br>
then open the terminal in the little-noro directory and run ./yoursystembinary<br>
### Windows:<br>
> ./Little_Noro.exe<br>


> [!NOTE]
> You don't need any dependancies for the binaries.

## Build Instructions:
You can download the source code and go to /Little_Noro/Main and install fyne framework and run:<br>
> fyne build main.go

> [!NOTE]
> Keep the images and sound files in the same directory as main.go or the binary file.





