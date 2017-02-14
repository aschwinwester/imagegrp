# imagegrp
Grap images from a HTML page or grap multiple images from one image URL using a simple algorithm

### usage
Usage of imagegrp:
  -cssClass string
        css class value
  -f string
        Folder name
  -v    Verbose logging
  

### the options  
  Option cssClass is optional. This option selects only images which contain in HTML a img element with the given class.
  Option f, the folder name is optional.
  Option v, verbose logging is optional.

### example
  ./imagegrp -v -f=/Users/aschwin/Pictures/dogs/ -cssClass=imagecache http://www.daws.org/dogs