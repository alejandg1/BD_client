package tui

func SetIcon(Db string) string {
	switch Db {
	case "postgres":
		return " "
	case "mysql":
		return " "
	case "sqlite":
		return " "
	default:
		return ""
	}
}

func SetCursor(cursor int, index int) string {
  if cursor == index {
    return "❯"
  }
  return " "
}
