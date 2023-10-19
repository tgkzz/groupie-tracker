let loc = document.getElementsByClassName('loc');
        console.log(loc);
        let arr = []
        for (key of loc) {
            arr.push(key.innerHTML.replace(/[^\w\s]|_/g, " "));
        }
        var map
        function myMap() {
            var geocoder = new google.maps.Geocoder();
            var latlng = new google.maps.LatLng(51.508742, -0.120850);
            var mapOptions = {
                zoom: 3,
                center: latlng
            }
            map = new google.maps.Map(document.getElementById('googleMap'), mapOptions);
            for (key of arr) {
                geocoder.geocode({
                    'address': key
                }, function (results, status) {
                    if (status == 'OK') {
                        map.setCenter(results[0].geometry.location);
                        let marker = new google.maps.Marker({
                            map: map,
                            animation: google.maps.Animation.BOUNCE,
                            position: results[0].geometry.location
                        });
                        const infowindow = new google.maps.InfoWindow({
                            content: "<p>Marker Location:" + marker.getPosition() + "</p>",
                        });
                        google.maps.event.addListener(marker, "click", () => {
                            infowindow.open(map, marker);
                        });
                        google.maps.event.addListener(marker, 'click', function () {
                            var pos = map.getZoom();
                            map.setZoom(9);
                            map.setCenter(marker.getPosition());
                            window.setTimeout(function () {
                                map.setZoom(pos);
                            }, 3000);
                        })
                    } else {
                        console.log("Don't work because:" + status)
                    }
                })
            }
        }