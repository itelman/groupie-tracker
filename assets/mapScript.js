function myMap() {
    var mapProp = {
        center: new google.maps.LatLng(51.508742, -0.120850),
        zoom: 1,
    };
    var map = new google.maps.Map(document.getElementById("googleMap"), mapProp);
    var apiKey = document.getElementById("api_key");

    let locations = document.getElementsByClassName('left');
    var infowindow = new google.maps.InfoWindow();

    for (var i = 0; i < locations.length; i++) {
        try {
            fetch('https://maps.googleapis.com/maps/api/geocode/json?key='+apiKey+'&address=' + locations[i].textContent)
                .then(response => response.json())
                .then(data => {
                    var marker = new google.maps.Marker({
                        position: new google.maps.LatLng(data.results[0].geometry.location.lat, data.results[0].geometry.location.lng),
                        map: map,
                    });

                    google.maps.event.addListener(marker, 'click', function () {
                        infowindow.setContent(data.results[0].formatted_address);
                        infowindow.open(map, this);
                    });
                })
        } catch (error) {
            console.log(error);
        }
    }
}
