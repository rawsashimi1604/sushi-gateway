import c3, { ChartConfiguration } from "c3";
import React, { useEffect, useRef } from "react";

interface ChartProps {
  config: ChartConfiguration;
}

function Chart({ config }: ChartProps) {
  const chartInstance = useRef<c3.ChartAPI | null>(null);
  const chartRef = useRef<HTMLDivElement>(null);

  // Update the chart when gateway changes, simply update the data not rerender the whole chart
  useEffect(() => {
    if (chartRef && chartRef.current) {
      if (chartInstance.current) {
        chartInstance.current.destroy();
        chartInstance.current = c3.generate({
          ...config,
          bindto: chartRef.current,
        });
      } else {
        // Only generate a new chart if there isn't one already
        chartInstance.current = c3.generate({
          ...config,
          bindto: chartRef.current,
        });
      }
    }
  }, [config]);

  // Destroy when chart unmounts, prevent memory leaks
  useEffect(() => {
    return () => {
      if (chartInstance.current) {
        chartInstance.current.destroy();
      }
    };
  }, []);

  return <div ref={chartRef}></div>;
}

export default Chart;
