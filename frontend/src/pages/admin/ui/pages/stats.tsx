import AppSidebar from "../components/sidebar";
import Header from "../components/header";
import Stats from "../components/stats";

const StatsPage: React.FC = () => {
  const data = {
    labels: ["Янв", "Фев", "Мар", "Апр", "Май"],
    sales: [500, 1000, 1500, 2000, 2500],
    visits: [1500, 2000, 2500, 3000, 3500],
  };

  return (
    <>
      <Header />
      <AppSidebar />
      <main>
        <Stats data={data} />
      </main>
    </>
  );
};

export default StatsPage;
