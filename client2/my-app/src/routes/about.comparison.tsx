import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/about/comparison')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>
    Deployment and  maintenance
    Self-managed and self-maintained
    Cloud hosted, automated, maintained and scaled by  the WireMock platform
    User Interface
    Managed via scripts and code
    Easy web UI and powerful local CLI
    Scalability
    Manually orchestrated and manually scaled
    Hosted with unlimited scalability
    Team collaboration
    Individual developer only
    Team collaboration and RBAC
    User Management
    Individual developer only
    RBAC, SSO
    K8s support
    No k8s-specific implementation
    K8s-native platform distribution
    Security
    No out of the box mock security
    Header-match, HTTP basic, OpenID Connect
    Protocols supported
    REST (gRPC / GraphQL via community extensions)
    REST, GraphQL, gRPC,  custom APIs
    OpenAPI / Swagger Importing
    Not natively supported
    Swagger, OpenAPI, WireMock OSS import
    Deployment and  maintenance
    Self-managed and self-maintained
    Cloud hosted, automated, maintained and scaled by  the WireMock platform
    Validator to alert in case of discrepancy between the mock and the API specification
    Not supported
    Drift detection with soft and hard validation
    Chaos Engineering Module
    Not supported
    Multiple chaos modes
    Mock-based API prototyping
    Not supported
    Collaborative interface to mock APIs that will generate open API spec and internal documentation portal
    Mock processes
    Running process per individual API mocked
    Hosted with unlimited scalability
    Performance metrics  dashboard
    No metrics available
    Visibility into usage and  performance
    Mock API templates (Public as well as private Mock templates library)
    No UI or templating
    Thousands of  public templates as well as ability to create company  private mock templates
    Support
    Community Support
    Enterprise Support including guaranteed SLAs
    GIT Integration
    Not supported
    Supported
    Import data from CSV files to be used by the mock
    Not supported
    Supported
    Stateful scenarios
    Not supported
    Supported
</div>
}
