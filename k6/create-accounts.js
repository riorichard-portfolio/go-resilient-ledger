import http from 'k6/http';
import { check } from 'k6';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js';

export const options = {
  scenarios: {
    serious_stress_test: {
      executor: 'ramping-arrival-rate',
      startRate: 50,             // Mulai dari 50 RPS
      timeUnit: '1s',
      preAllocatedVUs: 2000,     // Alokasi VU awal yang besar untuk lonjakan tinggi
      maxVUs: 20000,              // Batas maksimal VU (mengantisipasi lonjakan latensi)
      stages: [
        { duration: '30s', target: 100 },    // Warm-up ke 100 RPS
        { duration: '30s', target: 500 },    // Naik ke 500 RPS
        { duration: '30s', target: 1000 },   // Naik ke 1.000 RPS (1k)
        { duration: '1m', target: 5000 },   // Genjot ke 5.000 RPS selama 1 menit
        // { duration: '1m', target: 7000 },  //  ke 7.000 RPS selama 1 menit
        // { duration: '1m', target: 10000 },  // Peak/Puncak ke 10.000 RPS selama 1 menit
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

export default function () {
  const url = 'http://localhost:3000/accounts'; // sesuaikan port
  const payload = JSON.stringify({
    name: 'rio_' + __VU + '_' + __ITER,
    balance: 0,
  });
  const params = { headers: { 'Content-Type': 'application/json' } };

  const res = http.post(url, payload, params);
  check(res, {
    'success (200/201)': (r) => r.status === 200 || r.status === 201,
    'queue full (429)': (r) => r.status === 429,
  });
}

export function handleSummary(data) {
  return {
    'stdout': textSummary(data, { indent: ' ', enableColors: true }),
    'create.account.summary.json': JSON.stringify(data), // Menyimpan raw JSON hasil pengujian
  };
}