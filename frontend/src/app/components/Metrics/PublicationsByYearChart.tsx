import React, { useState } from 'react';
import {
  BarChart,
  Bar,
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from 'recharts';

type PublicationsByYearChartProps = {
  data: { year: number; count: number }[];
};

export default function PublicationsByYearChart({ data }: PublicationsByYearChartProps) {
  const [chartType, setChartType] = useState<'bar' | 'line'>('bar');

  return (
    <div className="bg-background rounded-lg border border-border-color p-6">
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-white text-lg font-semibold">Publicações por Ano</h3>
        <div className="inline-flex rounded-md bg-gray-700">
          {['bar', 'line'].map((type) => (
            <button
              key={type}
              onClick={() => setChartType(type as 'bar' | 'line')}
              className={`px-3 py-1 text-sm font-medium rounded-md focus:outline-none cursor-pointer ${
                chartType === type
                  ? 'bg-yellow-400 text-black'
                  : 'text-gray-400 hover:bg-gray-600'
              }`}
              type="button"
            >
              {type === 'bar' ? 'Barras' : 'Linhas'}
            </button>
          ))}
        </div>
      </div>
      <div className="h-80">
        <ResponsiveContainer width="100%" height="100%">
          {chartType === 'bar' ? (
            <BarChart
              data={data}
              margin={{ top: 10, right: 10, left: -20, bottom: 20 }}
            >
              <CartesianGrid strokeDasharray="3 3" stroke="#2D2D3A" />
              <XAxis
                dataKey="year"
                tick={{ fill: '#C4C7CA' }}
                tickLine={{ stroke: '#8E9196' }}
                axisLine={{ stroke: '#8E9196' }}
              />
              <YAxis
                tick={{ fill: '#C4C7CA' }}
                tickLine={{ stroke: '#8E9196' }}
                axisLine={{ stroke: '#8E9196' }}
              />
              <Tooltip
                contentStyle={{
                  backgroundColor: '#1E1E2E',
                  borderColor: '#8E9196',
                  color: '#FFFFFF',
                }}
                itemStyle={{ color: '#FFFFFF' }}
                cursor={{ fill: 'rgba(255, 196, 0, 0.1)' }}
              />
              <Bar dataKey="count" fill="#FFC400" name="Publicações" />
            </BarChart>
          ) : (
            <LineChart
              data={data}
              margin={{ top: 10, right: 10, left: -20, bottom: 20 }}
            >
              <CartesianGrid strokeDasharray="3 3" stroke="#2D2D3A" />
              <XAxis
                dataKey="year"
                tick={{ fill: '#C4C7CA' }}
                tickLine={{ stroke: '#8E9196' }}
                axisLine={{ stroke: '#8E9196' }}
              />
              <YAxis
                tick={{ fill: '#C4C7CA' }}
                tickLine={{ stroke: '#8E9196' }}
                axisLine={{ stroke: '#8E9196' }}
              />
              <Tooltip
                contentStyle={{
                  backgroundColor: '#1E1E2E',
                  borderColor: '#8E9196',
                  color: '#FFFFFF',
                }}
                itemStyle={{ color: '#FFFFFF' }}
                cursor={{ stroke: '#FFC400', strokeWidth: 2 }}
              />
              <Line
                type="monotone"
                dataKey="count"
                stroke="#FFC400"
                strokeWidth={2}
                dot={{ fill: '#FFC400', stroke: '#FFC400' }}
                activeDot={{ r: 6, fill: '#FFC400' }}
                name="Publicações"
              />
            </LineChart>
          )}
        </ResponsiveContainer>
      </div>
    </div>
  );
}