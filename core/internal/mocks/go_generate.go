package mocks

//go:generate mockery --quiet --srcpkg github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/flux_aggregator_wrapper --name FluxAggregatorInterface --output . --case=underscore --structname FluxAggregator --filename flux_aggregator.go
//go:generate mockery --quiet --srcpkg github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/flags_wrapper --name FlagsInterface --output . --case=underscore --structname Flags --filename flags.go
//go:generate mockery --quiet --srcpkg github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/aggregator_v3_interface --name AggregatorV3InterfaceInterface --output ../../services/vrf/mocks/ --case=underscore --structname AggregatorV3Interface --filename aggregator_v3_interface.go
//go:generate mockery --quiet --srcpkg github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_coordinator_v2 --name VRFCoordinatorV2Interface --output ../../services/vrf/mocks/ --case=underscore --structname VRFCoordinatorV2Interface --filename vrf_coordinator_v2.go
