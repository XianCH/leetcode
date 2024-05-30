import React, { useState, useEffect } from 'react';

const ExchangeRates = () => {
  const [exchangeRates, setExchangeRates] = useState({});
  const [loading, setLoading] = useState(true);
  const [baseCurrency, setBaseCurrency] = useState('CNY');
  const apiEndpoint = 'https://v6.exchangerate-api.com/v6/49ec4d9c4632b076728e35af/latest/';

  useEffect(() => {
    const fetchExchangeRates = async () => {
      try {
        const response = await fetch(`${apiEndpoint}${baseCurrency}`);
        const data = await response.json();
        setExchangeRates(data.conversion_rates);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching exchange rates:', error);
        setLoading(false);
      }
    };

    fetchExchangeRates();
  }, [baseCurrency]);

  const handleCurrencyChange = (e) => {
    setBaseCurrency(e.target.value);
    setLoading(true);
  };

  if (loading) {
    return <div className="text-center mt-16">Loading...</div>;
  }

  return (
    <div className="exchange-rates text-center mt-16">
      <h1 className="text-4xl font-bold mb-8">实时汇率</h1>
      <div className="mt-4">
        <label htmlFor="baseCurrency" className="mr-2">当前货币：</label>
        <select
          id="baseCurrency"
          value={baseCurrency}
          onChange={handleCurrencyChange}
          className="form-select p-2 border rounded"
        >
          <option value="CNY">CNY</option>
          <option value="USD">USD</option>
          <option value="JPY">JPY</option>
          <option value="EUR">EUR</option>
          <option value="GBP">GBP</option>
        </select>
      </div>
      <ul className="list-group mt-4">
        {Object.entries(exchangeRates).map(([currency, rate]) => (
          <li key={currency} className="list-group-item p-2 border-b">
            {currency}: {rate}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ExchangeRates;
