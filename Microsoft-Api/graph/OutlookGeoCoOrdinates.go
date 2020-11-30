package graph

import(

)

type GeoCoordinates struct{
  Accuracy         float64 `json "accuracy"`
  Altitude         float64 `json "altitude"`
  AltitudeAccuracy float64 `json "altitudeAccuracy"`
  Latitude         float64 `json "latitude"`
  Longitude        float64 `json "longitude"`
}
