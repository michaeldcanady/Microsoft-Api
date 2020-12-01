package graph

import(
  "time"
)

type Photo struct{
  CameraMake          string    `json "cameraMake"`
  CameraModel         string    `json "cameraModel"`
  ExposureDenominator float64   `json "exposureDenominator"`
  ExposureNumerator   float64   `json "exposureNumerator"`
  FNumber             float64   `json "fNumber"`
  FocalLength         float64   `json "focalLength"`
  Iso                 int       `json "iso"`
  Orientation         int       `json "orientation"`
  TakenDateTime       time.Time `json "takenDateTime"`
}
