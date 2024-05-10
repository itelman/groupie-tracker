//@ts-ignore
const { AdvancedMarkerElement, PinElement } = await google.maps.importLibrary("marker");
const { Map } = await google.maps.importLibrary("maps");

function myMap() {
    
    var mapProp = {
        center: new google.maps.LatLng(51.508742, -0.120850),
        zoom: 1,
        mapId: "d750397acb3eae28",
    };
    var map = new google.maps.Map(document.getElementById("googleMap"), mapProp);

    let locations = document.getElementsByClassName('left');
    var infowindow = new google.maps.InfoWindow();

    for (var i = 0; i < locations.length; i++) {
        try {
            fetch('https://maps.googleapis.com/maps/api/geocode/json?key=AIzaSyDhPphTq-AjtymkeraRqlz6RFz9NNWdZi4&address=' + locations[i].textContent)
                .then(response => response.json())
                .then(data => {
                    var marker = new google.maps.marker.AdvancedMarkerView({
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