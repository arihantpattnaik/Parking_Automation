// static/script.js
// Get all parking slots
async function getAllParkings() {
    try {
        const response = await fetch('/api/parkings');
        const data = await response.json();
        displayParkings(data);
    } catch (error) {
        console.error('Error:', error);
    }
}

// Display parking slots
function displayParkings(parkings) {
    const list = document.getElementById('parkingList');
    list.innerHTML = '';
    parkings.forEach(parking => {
        const slot = document.createElement('div');
        slot.className = 'parking-slot';
        slot.innerHTML = `
            <p><strong>Floor:</strong> ${parking.parkfloor}</p>
            <p><strong>Price:</strong> $${parking.parkprice}</p>
            <p><strong>Available:</strong> ${parking.parkavailable ? 'Yes' : 'No'}</p>
            <p><strong>Owner:</strong> ${parking.owner?.ownername || 'N/A'}</p>
            <button onclick="deleteParking('${parking._id}')" class="delete-btn">Delete</button>
            <button onclick="markUnavailable('${parking._id}')" class="unavailable-btn">
                Mark Unavailable
            </button>
        `;
        list.appendChild(slot);
    });
}

// Create new parking slot
document.getElementById('createForm').onsubmit = async (e) => {
    e.preventDefault();
    const parkingData = {
        parkfloor: parseInt(document.getElementById('parkFloor').value),
        parkprice: parseFloat(document.getElementById('parkPrice').value),
        parkavailable: true,
        owner: {
            ownername: document.getElementById('ownerName').value,
            ownernumber: document.getElementById('ownerNumber').value
        }
    };

    try {
        await fetch('/api/parking', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(parkingData),
        });
        getAllParkings();
        e.target.reset(); // Clear form after successful submission
    } catch (error) {
        console.error('Error:', error);
    }
};

// Delete parking slot
async function deleteParking(id) {
    if (confirm('Are you sure you want to delete this parking slot?')) {
        try {
            await fetch(`/api/parking/${id}`, {
                method: 'DELETE',
            });
            getAllParkings();
        } catch (error) {
            console.error('Error:', error);
        }
    }
}

// Mark parking slot as unavailable
async function markUnavailable(id) {
    try {
        await fetch(`/api/parking/${id}`, {
            method: 'PUT',
        });
        getAllParkings();
    } catch (error) {
        console.error('Error:', error);
    }
}

// Load parking slots on page load
document.addEventListener('DOMContentLoaded', getAllParkings);