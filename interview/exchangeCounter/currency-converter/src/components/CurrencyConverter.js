import React, { useState, useEffect } from 'react';

const CurrencyConverter = () => {
  const [amount, setAmount] = useState('');
  const [convertedAmount, setConvertedAmount] = useState('');
  const [exchangeRates, setExchangeRates] = useState({});
  const [loading, setLoading] = useState(true);
  const [sourceCurrency, setSourceCurrency] = useState('USD');
  const [targetCurrency, setTargetCurrency] = useState('CNY');

  const apiEndpoint = 'https://v6.exchangerate-api.com/v6/49ec4d9c4632b076728e35af/latest/';

  useEffect(() => {
    const fetchExchangeRates = async () => {
      try {
        const response = await fetch(`${apiEndpoint}${sourceCurrency}`);
        const data = await response.json();
        setExchangeRates(data.conversion_rates);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching exchange rates:', error);
        setLoading(false);
      }
    };

    fetchExchangeRates();
  }, [sourceCurrency]);

  const handleConvert = () => {
    if (!amount || !exchangeRates[targetCurrency]) return;
    const result = (amount * exchangeRates[targetCurrency]).toFixed(2);
    setConvertedAmount(result);
  };

  const handleSourceCurrencyChange = (e) => {
    setSourceCurrency(e.target.value);
    setLoading(true);
  };

  const handleTargetCurrencyChange = (e) => {
    setTargetCurrency(e.target.value);
  };

  if (loading) {
    return <div className="text-center mt-16">Loading...</div>;
  }

  return (
    <div className="currency-converter text-center mt-16">
      <h1 className="text-4xl font-bold mb-8">汇率转换器</h1>
      <div className="mb-4">
        <label className="mr-2">源货币：</label>
        <select
          value={sourceCurrency}
          onChange={handleSourceCurrencyChange}
          className="form-select p-2 border rounded mr-4"
        >
          {Object.keys(exchangeRates).map((currency) => (
            <option key={currency} value={currency}>{currency}</option>
          ))}
        </select>
        <input
          type="number"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          placeholder={`Amount in ${sourceCurrency}`}
          className="form-input p-2 border rounded"
        />
      </div>
      <div className="mb-4">
        <label className="mr-2">目标货币：</label>
        <select
          value={targetCurrency}
          onChange={handleTargetCurrencyChange}
          className="form-select p-2 border rounded mr-4"
        >
          {Object.keys(exchangeRates).map((currency) => (
            <option key={currency} value={currency}>{currency}</option>
          ))}
        </select>
        <button onClick={handleConvert} className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700">Convert</button>
      </div>
      {convertedAmount && (
        <div className="mt-4">
          <p>Converted Amount: {convertedAmount} {targetCurrency}</p>
        </div>
      )}
    </div>
  );
};

export default CurrencyConverter;
