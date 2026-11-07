import http from 'k6/http';
import { check } from 'k6';

export const options = {
    scenarios: {
        serious_stress_test: {
            executor: 'ramping-arrival-rate',
            startRate: 50,             // Mulai dari 50 RPS
            timeUnit: '1s',
            preAllocatedVUs: 2000,     // Alokasi VU awal yang besar untuk lonjakan tinggi
            maxVUs: 8000,              // Batas maksimal VU (mengantisipasi lonjakan latensi)
            stages: [
                { duration: '30s', target: 100 },    // Warm-up ke 100 RPS
                { duration: '30s', target: 500 },    // Naik ke 500 RPS
                { duration: '30s', target: 1000 },   // Naik ke 1.000 RPS (1k)
                { duration: '1m', target: 5000 },   // Genjot ke 5.000 RPS selama 1 menit
                { duration: '1m', target: 10000 },  // Peak/Puncak ke 10.000 RPS selama 1 menit
                { duration: '30s', target: 0 },      // Ramp-down ke 0 RPS (Cool down)
            ],
        },
    },
    thresholds: {
        // Threshold lebih longgar karena beban ekstrem 10k RPS
        http_req_duration: ['p(95)<1000', 'p(99)<2000'],
        http_req_failed: ['rate<0.05'], // Toleransi error maksimal 5% di beban puncak
    },
};

// Daftar sampel nama untuk dikombinasikan secara random
const baseNames = ['rio', 'budi', 'siti', 'joko', 'ani', 'eko', 'dewi', 'agus', 'rudi', 'siska'];

export default function () {
    const url = 'http://localhost:3000/accounts';

    // 1. Generate random name (Kombinasi nama dasar + angka unik agar tidak duplikat persis)
    const randomNamePrefix = baseNames[Math.floor(Math.random() * baseNames.length)];
    const uniqueSuffix = Math.floor(Math.random() * 100000);
    const name = `${randomNamePrefix}_${uniqueSuffix}`;

    // 2. Generate random balance >= 0 (Contoh rentang 0 sampai 10.000.000)
    const balance = Math.floor(Math.random() * 10000000);

    const payload = JSON.stringify({
        name: name,
        balance: balance,
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const res = http.post(url, payload, params);

    check(res, {
        'status is 200/201': (r) => r.status === 200 || r.status === 201,
    });
}

// Menampilkan persentil di akhir test
export function handleSummary(data) {
    const duration = data.metrics.http_req_duration.values;

    console.log('\n========================================');
    console.log('       CUSTOM LATENCY PERCENTILES       ');
    console.log('========================================');
    console.log(`> P50 (Median) : ${duration['p(50)'].toFixed(2)} ms`);
    console.log(`> P95          : ${duration['p(95)'].toFixed(2)} ms`);
    console.log(`> P99          : ${duration['p(99)'].toFixed(2)} ms`);
    console.log(`> Average      : ${duration['avg'].toFixed(2)} ms`);
    console.log(`> Max          : ${duration['max'].toFixed(2)} ms`);
    console.log('========================================\n');

    return {};
}