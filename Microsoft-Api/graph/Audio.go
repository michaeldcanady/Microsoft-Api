package graph

import(

)

type Audio struct{
  Album             string `json "album"`
  AlbumArtist       string `json "albumArtist"`
  Artist            string `json "artist"`
  Bitrate           int    `json "bitrate"`
  Composers         string `json "composers"`
  Copyright         string `json "copyright"`
  Disc              int    `json "disc"`
  DiscCount         int    `json "discCount"`
  Duration          int    `json "duration"`
  Genre             string `json "genre"`
  HasDrm            bool   `json "hasDrm"`
  IsVariableBitrate bool   `json "isVariableBitrate"`
  Title             string `json "title"`
  Track             int    `json "track"`
  TrackCount        int    `json "trackCount"`
  Year              int    `json "year"`
}
