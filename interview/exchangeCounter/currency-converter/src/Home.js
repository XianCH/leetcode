import React from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
  return (
    <div className="text-center mt-16">
      <h1 className="text-4xl font-bold mb-8">汇率计算器</h1>
      <div className="space-x-4">
        <Link className="px-6 py-2 bg-blue-500 text-white rounded hover:bg-blue-700" to="/converter">汇率转换器</Link>
        <Link className="px-6 py-2 bg-gray-500 text-white rounded hover:bg-gray-700" to="/rates">实时汇率</Link>
      </div>
    </div>
  );
};

export default Home;