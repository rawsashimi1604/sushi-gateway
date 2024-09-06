import { IoMdInformationCircle } from "react-icons/io";
import Header from "../../components/typography/Header";
import DashboardCard from "../../components/layout/DashboardCard";

function Summary() {
  return (
    <DashboardCard className="flex flex-col gap-2 p-6 ">
      <Header text="summary" align="left" size="md" />
      <div className="flex items-center gap-8">
        {/* Services */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">services</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header text="2" align="left" size="lg" />
        </div>

        {/* Routes */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">routes</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header text="8" align="left" size="lg" />
        </div>

        {/* Plugins */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">plugins</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header text="15" align="left" size="lg" />
        </div>

        {/* Last Updated */}
        <div className="flex flex-col gap-2 items-start justify-start">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">last updated</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header text="24 Sept 2024 15:00:00" align="left" size="lg" />
        </div>
      </div>
    </DashboardCard>
  );
}

export default Summary;
