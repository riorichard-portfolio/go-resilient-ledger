const { performance } = require('perf_hooks');

// Helper untuk nama acak
function getRandomName() {
    const firstNames = ['Budi', 'Ani', 'Siti', 'Joko', 'Dewi', 'Eko', 'Rina', 'Agus', 'Bambang', 'Siska', 'Rio', 'Dedi'];
    const lastNames = ['Santoso', 'Pratama', 'Wijaya', 'Kusuma', 'Putri', 'Lestari', 'Hidayat', 'Nugroho', 'Siregar', 'Utami'];
    const randomFirst = firstNames[Math.floor(Math.random() * firstNames.length)];
    const randomLast = lastNames[Math.floor(Math.random() * lastNames.length)];
    const randomNum = Math.floor(Math.random() * 9000) + 1000;
    return `${randomFirst} ${randomLast} ${randomNum}`;
}

// Helper untuk balance acak
function getRandomBalance() {
    return Math.floor(Math.random() * 10000000);
}

// Helper untuk menghitung persentil latensi
function calculatePercentile(arr, percentile) {
    if (arr.length === 0) return 0;
    const sorted = [...arr].sort((a, b) => a - b);
    const index = Math.ceil((percentile / 100) * sorted.length) - 1;
    return sorted[Math.max(0, index)];
}

async function sendRpsWithMetrics(rps, durationInSeconds) {
    const url = 'http://localhost:3000/accounts';
    const headers = { 'Content-Type': 'application/json' };

    console.log(`Menyiapkan pre-prepared data untuk total ${rps * durationInSeconds} request...`);

    // 1. Pre-prepared data
    const allBatches = [];
    for (let sec = 0; sec < durationInSeconds; sec++) {
        const batch = [];
        for (let i = 0; i < rps; i++) {
            batch.push(JSON.stringify({
                name: getRandomName(),
                balance: getRandomBalance()
            }));
        }
        allBatches.push(batch);
    }

    console.log(`Data siap! Memulai load test: ${rps} RPS selama ${durationInSeconds} detik...\n`);

    let totalSuccess = 0;
    let totalFailed = 0;
    const latencies = []; // Menyimpan waktu durasi tiap request (dalam milidetik)

    const testStartTime = performance.now();

    // 2. Eksekusi test per detik
    for (let sec = 0; sec < durationInSeconds; sec++) {
        const secStartTime = performance.now();
        const currentBatch = allBatches[sec];
        const requests = [];

        for (let i = 0; i < rps; i++) {
            const reqStart = performance.now(); // Catat waktu mulai request
            
            const req = fetch(url, {
                method: 'POST',
                headers: headers,
                body: currentBatch[i]
            })
            .then(res => {
                const reqDuration = performance.now() - reqStart; // Hitung durasi
                latencies.push(reqDuration);
                
                if (res.ok) {
                    totalSuccess++;
                } else {
                    totalFailed++;
                }
            })
            .catch(() => {
                const reqDuration = performance.now() - reqStart;
                latencies.push(reqDuration);
                totalFailed++;
            });

            requests.push(req);
        }

        await Promise.all(requests);

        const elapsed = performance.now() - secStartTime;
        if (elapsed < 1000) {
            await new Promise(resolve => setTimeout(resolve, 1000 - elapsed));
        }

        console.log(`Detik ke-${sec + 1} selesai.`);
    }

    const testTotalDuration = (performance.now() - testStartTime) / 1000; // Konversi ke detik
    const totalRequests = totalSuccess + totalFailed;
    
    // 3. Kalkulasi Metrik Akhir
    const actualRps = (totalRequests / testTotalDuration).toFixed(2);
    const p50 = calculatePercentile(latencies, 50).toFixed(2);
    const p95 = calculatePercentile(latencies, 95).toFixed(2);
    const p99 = calculatePercentile(latencies, 99).toFixed(2);
    const avgLatency = (latencies.reduce((a, b) => a + b, 0) / (latencies.length || 1)).toFixed(2);

    // 4. Cetak Report
    console.log(`\n========================================`);
    console.log(`         HASIL LAPORAN LOAD TEST        `);
    console.log(`========================================`);
    console.log(`Total Request    : ${totalRequests}`);
    console.log(`Total Berhasil   : ${totalSuccess}`);
    console.log(`Total Gagal      : ${totalFailed}`);
    console.log(`Durasi Pengujian : ${testTotalDuration.toFixed(2)} detik`);
    console.log(`----------------------------------------`);
    console.log(`Throughput Aktual: ${actualRps} RPS`);
    console.log(`----------------------------------------`);
    console.log(`Latency Rata-rata: ${avgLatency} ms`);
    console.log(`Latency P50      : ${p50} ms`);
    console.log(`Latency P95      : ${p95} ms`);
    console.log(`Latency P99      : ${p99} ms`);
    console.log(`========================================`);
}

// Contoh pemanggilan: Target 100 RPS selama 5 detik
sendRpsWithMetrics(10000, 1);