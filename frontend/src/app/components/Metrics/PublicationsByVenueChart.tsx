// components/Metrics/PublicationsByVenueChart.tsx

import React, { useMemo } from "react";
import { PieChart, Pie, Cell, Tooltip, Legend, ResponsiveContainer } from "recharts";

type VenueMetric = {
  name: string;
  count: number;
};

type PublicationsByVenueChartProps = {
  data: VenueMetric[];
  /** 
   * Número máximo de fatias “individuais” a exibir.
   * Os demais serão agrupados em “Outros”.
   */
  topN?: number;
  colors?: string[];
};

export default function PublicationsByVenueChart({
  data,
  topN = 6, // exibe as 6 maiores fatias e agrupa o resto
  colors = ["#FFC400", "#FF8042", "#00C49F", "#0088FE", "#FFBB28", "#8884d8"],
}: PublicationsByVenueChartProps) {
  
  // 1) Prepara os dados: ordena por count descendente, pega topN e soma o restante como “Outros”
  const processedData = useMemo(() => {
    if (!data || data.length === 0) return [];

    // 1.1) Ordena por count descendente
    const sorted = [...data].sort((a, b) => b.count - a.count);

    // 1.2) Se tivermos menos ou exatamente topN itens, não precisa agrupar nada
    if (sorted.length <= topN) {
      return sorted;
    }

    // 1.3) Fatiar os topN
    const topItems = sorted.slice(0, topN);

    // 1.4) Soma todos os counts restantes
    const restSum = sorted
      .slice(topN)
      .reduce((acc, item) => acc + item.count, 0);

    // 1.5) Cria um registro “Outros” apenas se houver algo fora do topN
    if (restSum > 0) {
      topItems.push({
        name: "Outros",
        count: restSum,
      });
    }

    return topItems;
  }, [data, topN]);

  // 2) Estima o total (para calcular percentagens, legendas etc.)
  const totalCount = useMemo(() => {
    return processedData.reduce((sum, entry) => sum + entry.count, 0);
  }, [processedData]);

  return (
    <div className="bg-background rounded-lg border border-border-color p-6">
      <h3 className="text-white text-lg font-semibold mb-4">
        Distribuição por Periódico/Conferência
      </h3>
      <div className="h-80">
        <ResponsiveContainer width="100%" height="100%">
          <PieChart>
            <Pie
              data={processedData}
              cx="50%"
              cy="50%"
              dataKey="count"
              nameKey="name"
              // remove as linhas que conectam label à fatia (fica mais limpo)
              labelLine={false}
              // exibe apenas o percentual para fatias ≥ 5%, senão exibe string vazia
              label={({ index, percent }) => {
                // se a fatia ocupar menos de 5% do total, não mostra label
                return percent * 100 >= 5
                  ? `${processedData[index].name} (${(percent * 100).toFixed(0)}%)`
                  : "";
              }}
              outerRadius="70%"
              fill="#8884d8"
            >
              {processedData.map((entry, index) => (
                <Cell
                  key={`cell-${index}`}
                  fill={colors[index % colors.length]}
                />
              ))}
            </Pie>

            <Tooltip
              formatter={(value: number, name: string, props: any) => [
                `${value} publicação${value > 1 ? "s" : ""}`,
                props.payload.name,
              ]}
              itemStyle={{ color: "#000" }}
              contentStyle={{
                backgroundColor: "#fff",
                borderColor: "#8E9196",
                borderRadius: "8px",
                padding: "8px",
                color: "#000",
              }}
            />

            <Legend
              formatter={(value: string) => (
                <span style={{ color: "#C4C7CA", fontSize: "0.875rem" }}>
                  {value}
                </span>
              )}
              layout="horizontal"
              verticalAlign="bottom"
              wrapperStyle={{ paddingTop: "0.5rem" }}
            />
          </PieChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}
