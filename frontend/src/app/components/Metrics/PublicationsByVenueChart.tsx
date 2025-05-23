import React from 'react';
import { PieChart, Pie, Cell, Tooltip, Legend, ResponsiveContainer } from 'recharts';

type PublicationsByVenueChartProps = {
  data: { name: string; count: number }[];
  colors?: string[];
};

export default function PublicationsByVenueChart({
  data,
  colors = ['#FFC400', '#FF8042', '#00C49F', '#0088FE', '#FFBB28', '#8884d8'],
}: PublicationsByVenueChartProps) {
  return (
    <div className="bg-background rounded-lg border border-border-color p-6">
      <h3 className="text-white text-lg font-semibold mb-4">
        Distribuição por Periódico/Conferência
      </h3>
      <div className="h-80">
        <ResponsiveContainer width="100%" height="100%">
          <PieChart>
            <Pie
              data={data}
              cx="50%"
              cy="50%"
              labelLine={false}
              label={({ name, percent }) => `${name} (${(percent * 100).toFixed(0)}%)`}
              outerRadius="70%"
              fill="#8884d8"
              dataKey="count"
              nameKey="name"
            >
              {data.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={colors[index % colors.length]} />
              ))}
            </Pie>
            <Tooltip
              contentStyle={{ backgroundColor: '#fefefe', borderColor: '#8E9196', color: '#FFFFFF' }}
              formatter={(value) => [`${value} publicações`]}
            />
            <Legend
              formatter={(value) => <span style={{ color: '#C4C7CA' }}>{value}</span>}
              layout="horizontal"
              verticalAlign="bottom"
            />
          </PieChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}