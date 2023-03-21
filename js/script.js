// Copyright (c) 2023 Mariam Kasim All rights reserved
//
// Created by: Mariam Kasim
// Created on: March 2023
// This file contains the JS functions for index.html

function enterClicked() {
  // input
  const streetname = document.getElementById("street-name").value
  const streetnumber = document.getElementById("number-entered").value

  // output
  document.getElementById("address").innerHTML =
    "Your address is: " + streetnumber + ", " + streetname + "."
}
