async function loadCSV(url) {
    const response = await fetch(url); 
    const data = await response.text(); 

    const rows = data.split('\n').map(row => row.split(','));

    // Map the rows into a format suitable for visualization
    const pcaData = rows.slice(1).map(row => ({
        // Principal Component 1
        x: parseFloat(row[0]), 
        // Principal Component 2
        y: parseFloat(row[1])  
    }));

    return pcaData;
}

async function createChart() {
    const pcaData = await loadCSV('../data/reduced_data.csv'); 

    const ctx = document.getElementById('pcaChart').getContext('2d');
    new Chart(ctx, {
        type: 'scatter',
        data: {
            datasets: [{
                label: 'PCA Data Visualization',
                data: pcaData, 
                backgroundColor: 'rgba(75, 192, 192, 0.6)'
            }]
        },
        options: {
            scales: {
                x: {
                    title: {
                        display: true,
                        text: 'Principal Component 1'
                    }
                },
                y: {
                    title: {
                        display: true,
                        text: 'Principal Component 2'
                    }
                }
            }
        }
    });
}

document.addEventListener('DOMContentLoaded', createChart);
