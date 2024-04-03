package gen

import (
	// "fmt"
	"net/url"
	"strings"

	"github.com/clark-john/gfontsgen/config"
	"github.com/clark-john/gfontsgen/json"
	"github.com/fatih/color"
	"github.com/samber/lo"
)

const FONT_URL = "https://fonts.googleapis.com/css"

func GenerateUrl(font json.FontItem, options *GenerateUrlOptions) (string, bool) {
  u := GetUrl()
  var warnings int

  var family strings.Builder

  family.WriteString(SpaceToPlus(font.Family))  
  family.WriteString(":")

  if !IsAll(options.Variants) {
    for i, variant := range options.Variants {
      family.WriteString(variant)

      /* error */
      v := ReplaceIWithItalic(variant)

      if !HasKey(font.Files, v) {
        color.HiYellow(
          "Warning: variant %s doesn't exist on font %s.", 
          v, font.Family,
        )
        warnings++
      }

      if i < len(options.Variants) - 1 {
        family.WriteString(",")
      }
    }
  } else {
    family.WriteString(strings.Join(lo.Map(font.Variants, func(item string, i int) string {
      if item != "italic" {
        return strings.ReplaceAll(item, "italic", "i")
      } else {
        return "italic"
      }
    }),","))
  }

  q := u.Query()

  q.Set("family", family.String())
  q.Set("display", "swap")

  u.RawQuery = q.Encode()

  return UnescapeQuery(u.String()), warnings > 0
}

func GenerateUrlMultiple(items []json.FontItem, options []config.OptionItem) (string, bool) {
  u := GetUrl()
  var warnings int
  fontMap := make(map[string][]string)

  var familyQuery strings.Builder
    
  for _, op := range options {
    item, isFound := lo.Find(items, func(item json.FontItem) bool {
      return strings.EqualFold(item.Family, op.FontFamily)
    })
    if !isFound {
      color.HiYellow(`Warning: Font family "%s" doesn't exist.`, op.FontFamily)
      warnings++
    } else {
      for _, variant := range op.Variants {
        if !HasKey(item.Files, ReplaceIWithItalic(variant)) {
          color.HiYellow(
            "Warning: variant %s doesn't exist on font %s.",
            variant,
            item.Family,
          )
          warnings++
        }
      }
      if warnings == 0 {
        fontMap[item.Family] = op.Variants
      }
    }
  }

  var index int
  for family, vars := range fontMap {
    familyQuery.WriteString(SpaceToPlus(family) + ":")
    familyQuery.WriteString(strings.Join(vars, ","))
    familyQuery.WriteString(func() string {
      if index == len(fontMap) - 1 {
        return ""
      } else {
        return "|"
      }
    }())
    index++
  }
  q := u.Query()

  q.Set("family", familyQuery.String())
  q.Set("display", "swap")

  u.RawQuery = q.Encode()

  return UnescapeQuery(u.String()), warnings > 0
}

func GetUrl() *url.URL {
  u, _ := url.Parse(FONT_URL)
  return u
}

func SpaceToPlus(text string) string {
  return strings.ReplaceAll(text, " ", "+")
}

func UnescapeQuery(str string) string {
  s, _ := url.QueryUnescape(str)
  return s
}

func IsAll(vars []string) bool {
  _, isFound := lo.Find(vars, func(item string) bool {
    return item == "all"
  })
  return isFound
}

func ReplaceIWithItalic(str string) string {
	suffix := str[len(str) - 1:]
	trimmed := str[:len(str) - 1]

	if suffix == "i" {
		return trimmed + "italic"
	} else {
		return str
	}
}

func HasKey(m map[string]string, value string) bool {
  for key := range m {
    if key == value {
      return true
    }
  }
  return false
}
