## Useful resources
[css-tricks](https://css-tricks.com/svg-sprites-use-better-icon-fonts/)

## Goal
- you have a number of svg icons in a folder
- you run `svg-coke src-dir dest-dir`

You get an output like this:
```html
<svg>
  <defs>
    <symbol id="icon-home" viewBox="0 0 1024 1024">
    	<title>home</title>
    	<path class="path1" d="M1024 590.444l-512-397.426-512 397.428v-162.038l512-397.426 512 397.428zM896 576v384h-256v-256h-256v256h-256v-384l384-288z"></path>
    </symbol>

    <symbol id="icon-home2" viewBox="0 0 1024 1024">
    	<title>home2</title>
    	<path class="path1" d="M512 32l-512 512 96 96 96-96v416h256v-192h128v192h256v-416l96 96 96-96-512-512zM512 448c-35.346 0-64-28.654-64-64s28.654-64 64-64c35.346 0 64 28.654 64 64s-28.654 64-64 64z"></path>
    </symbol>
  </defs>
</svg>
```

Then you can use it like this:
```css
.icon {
  display: inline-block;
  width: 2rem;
  height: 2rem;
  fill: currentColor;
}
```

```html
<svg class="icon icon-home">
  <use xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="#icon-home"></use>
</svg>
```

## todo
- add example html file with all the icons displayed
- imporve logging
- imporve error handling
