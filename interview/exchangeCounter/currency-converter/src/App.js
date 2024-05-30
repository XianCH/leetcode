import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Home from './Home';
import CurrencyConverter from './components/CurrencyConverter';
import ExchangeRates from './components/ExchangeRates';

function App() {
  return (
    <Router>
      <div className="container mx-auto p-4">
        <nav className="bg-gray-100 p-4 rounded shadow">
          <div className="flex justify-between items-center">
            <Link className="text-xl font-bold" to="/">汇率计算器</Link>
            <div>
              <Link className="px-4 py-2 text-blue-500 hover:underline" to="/converter">汇率转换器</Link>
              <Link className="px-4 py-2 text-blue-500 hover:underline" to="/rates">实时汇率</Link>
            </div>
          </div>
        </nav>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/converter" element={<CurrencyConverter />} />
          <Route path="/rates" element={<ExchangeRates />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
