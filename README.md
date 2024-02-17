# gfontsgen

### Google Fonts URL and Font Files generator

### Simple Usage:

List generation

```bash
# limit list to 10 font items
gfontsgen list -l 10
# search entire list with "com"
gfontsgen list -s com
# generate list of san-serif fonts
gfontsgen list -d sans-serif
# generate list of display fonts
gfontsgen list -c display
# write list to a file
gfontsgen list --write-to-file # default: fonts.json
```

Url generation

```bash
# generate css url
gfontsgen gen url montserrat # default: regular font only
# generate css url with specific weights/variants
gfontsgen gen url montserrat -v 200,regular,600,700i # extralight, regular, semibold, bold italic
# generate css url and copy it to clipboard
gfontsgen gen url montserrat -v 500i --copy
# generate css url with import url
gfontsgen gen url montserrat -v italic -i
# generate css url with all variants
gfontsgen gen url montserrat -v all
```

Font files generation

```bash
# generate font files to a folder (default: fonts) (default variants: regular only)
gfontsgen gen fontfiles montserrat
# generate font files with specific variants
gfontsgen gen fontfiles montserrat -v 200,300,600i # 200 regular, 300 regular, 600 italic
# generate font files with all variants
gfontsgen gen fontfiles montserrat -v all
# generate font files to a custom folder
gfontsgen gen fontfiles montserrat -p myfonts
# generate woff font files
gfontsgen gen fontfiles montserrat --woff
```
