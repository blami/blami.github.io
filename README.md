# My Personal Website
This repository contains my personal website and blog https://www.blami.net/.
It is hosted here on Github Pages and powered by Jekyll.

This README is mostly collection of my notes made to get things going locally
and solve problems.


## Development

### Setup
To setup environment for both blogging and developing blog following is needed:
``` shell
git clone
apt install zlib1g-dev
bundle install --path vendor/bundle
```

### Templates
Page templates are in usual location `_layouts/`. Most of them inherit
`_default` which is main for both __condensed__ and __non-condensed__ pages.

### Navigation
Navigation menu is handled automatically by looking at all pages that have
following properties:

- `group: nav` - page must be grouped in nav
- `menu: <title>` - page `<title>` will be shown in navbar
- `order: <n>` - pages in nav will be sorted using `<n>`

NOTE: Hamburger is kinda weirdly placed but I like it that way.

### Themes
Themes are implemented as bare CSS style color changes (they are color schemes
in fact) and all live in `_css/themes.css`. Theme toggling happens by setting
`theme-<name>` class setting on body element. This is persisted in
`localStorage` locally on client.
