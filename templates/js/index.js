$(document).ready(function() {
    let allLocations = [];
    
    for(let i = 1; i <= 52; i++) {
        $.ajax({
            url: '/location/' + i,  
            type: 'GET',
            dataType: 'json',
            success: function(data) {
                allLocations = allLocations.concat(data.locations);
                
                if(i === 52) {
                    let uniqueLocations = [...new Set(allLocations)];

                    uniqueLocations.forEach(function(location) {
                        $('#locationFilter').append($('<option>', {
                            value: location,
                            text: location.replace(/_/g, ' ')
                        }));
                    });
                }
            },
            error: function(error) {
                console.error("Ошибка при загрузке локации для группы " + i, error);
            }
        });
    }
});
