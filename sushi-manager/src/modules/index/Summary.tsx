import { IoMdInformationCircle } from "react-icons/io";
import Header from "../../components/typography/Header";

function Summary() {
  return (
    <div className="shadow-md flex flex-col gap-2 bg-white p-6 rounded-lg">
      <Header text="summary" align="left" size="md" />
      <div className="flex items-center gap-8">
        {/* Services */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3">
            <Header text="services" size="xs" align="left" />
            <IoMdInformationCircle className="text-xl mt-1" />
          </div>
          <Header text="2" align="left" size="xxl" />
        </div>

        {/* Routes */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3">
            <Header text="routes" size="xs" align="left" />
            <IoMdInformationCircle className="text-xl mt-1" />
          </div>
          <Header text="8" align="left" size="xxl" />
        </div>

        {/* Plugins */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3">
            <Header text="plugins" size="xs" align="left" />
            <IoMdInformationCircle className="text-xl mt-1" />
          </div>
          <Header text="15" align="left" size="xxl" />
        </div>

        {/* Last Updated */}
        <div className="flex flex-col gap-2 ">
          <div className="flex gap-3">
            <Header text="last updated" size="xs" align="left" />
            <IoMdInformationCircle className="text-xl mt-1" />
          </div>
          <Header text="24 Sept 2024 15:00:00" align="left" size="lg" />
        </div>
      </div>
    </div>
  );
}

export default Summary;
