# logger

log.SetFlags(log.LstdFlags | log.Lshortfile)
l := logger.Logger{
  Directory: "logssss",
  Prefix:    "feed_server",
}
l.InitStandardLogger(logger.FILE)
