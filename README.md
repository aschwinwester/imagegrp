# imagegrp
Grap images from a HTML page or grap multiple images from one image URL using a simple algorithm

[![Build Status](https://travis-ci.org/aschwinwester/imagegrp.svg?branch=master)](https://travis-ci.org/aschwinwester/imagegrp)


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
  imagegrp -v -f=/Users/aschwin/Pictures/dogs/ -cssClass=imagecache http://www.daws.org/dogs
  

### TODO
Should retrieve images of a page based on URL of single image.
So http://<website>/images/dog01.png are argument should retrieve dog02.png and more as well.