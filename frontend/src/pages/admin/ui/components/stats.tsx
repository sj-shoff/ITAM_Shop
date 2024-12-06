import { Card } from "@nextui-org/react";
import { Bar } from "react-chartjs-2";
import { StatsData } from "@pages/admin/index";

interface StatsProps {
  data: StatsData;
}

const Stats: React.FC<StatsProps> = ({ data }) => {
  const chartData = {
    labels: data.labels,
    datasets: [
      {
        label: "Продажи",
        data: data.sales,
        backgroundColor: "rgba(75, 192, 192, 0.6)",
      },
      {
        label: "Посещения",
        data: data.visits,
        backgroundColor: "rgba(153, 102, 255, 0.6)",
      }
    ],
  };

  return (
    <Card>
      <Bar data={chartData} />
    </Card>
  );
};

export default Stats;
