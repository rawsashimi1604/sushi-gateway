import Header from '../../components/typography/Header'
import ServiceDropdown from '../../components/sushi-gateway/ServiceDropdown'

function Services() {
    return (
        <div>
            <Header text="Services" align="left" size="md" />
            <div className='flex flex-col gap-3'>
                <ServiceDropdown />
                <ServiceDropdown />
                <ServiceDropdown />
            </div>
        </div>
    )
}

export default Services