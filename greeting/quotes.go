package greeting

func Hello(lang string) string {
    var quote string
    switch lang {
    case "de":
        quote = "Guten Tag"
    case "en":
        quote = "Hello"
    case "fr":
        quote = "Bonjour"
    case "ch":
        quote = "Grüzi"
    case "at":
        quote = "Servus"
    }
    return quote
}