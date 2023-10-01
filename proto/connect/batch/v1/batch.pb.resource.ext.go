package batchv1

// Key returns the key of the resource.
func (x ParsedProjectName) Key() string {
	return x.ProjectID
}

// String returns the string representation.
func (x ParsedProjectName) String() string {
	return x.Name()
}

// Key returns the key of the resource.
func (x ParsedLocationName) Key() string {
	return x.LocationID
}

// Parent returns the parent resource.
func (x ParsedLocationName) Parent() ParsedProjectName {
	return ParsedProjectName{
		ProjectID: x.ProjectID,
	}
}

// String returns the string representation.
func (x ParsedLocationName) String() string {
	return x.Name()
}

// Key returns the key of the resource.
func (x ParsedRegionName) Key() string {
	return x.RegionID
}

// Parent returns the parent resource.
func (x ParsedRegionName) Parent() ParsedProjectName {
	return ParsedProjectName{
		ProjectID: x.ProjectID,
	}
}

// String returns the string representation.
func (x ParsedRegionName) String() string {
	return x.Name()
}

// Key returns the key of the resource.
func (x ParsedNetworkName) Key() string {
	return x.NetworkID
}

// Parent returns the parent resource.
func (x ParsedNetworkName) Parent() ParsedProjectName {
	return ParsedProjectName{
		ProjectID: x.ProjectID,
	}
}

// String returns the string representation.
func (x ParsedNetworkName) String() string {
	return x.Name()
}

// Key returns the key of the resource.
func (x ParsedSubnetworkName) Key() string {
	return x.SubnetworkID
}

// String returns the string representation.
func (x ParsedSubnetworkName) String() string {
	return x.Name()
}

// Key returns the key of the resource.
func (x ParsedJobName) Key() string {
	return x.JobID
}

// Parent returns the parent resource.
func (x ParsedJobName) Parent() ParsedLocationName {
	return ParsedLocationName{
		ProjectID:  x.ProjectID,
		LocationID: x.LocationID,
	}
}

// String returns the string representation.
func (x ParsedJobName) String() string {
	return x.Name()
}

// Key returns the key of the resource.
func (x ParsedTaskGroupName) Key() string {
	return x.TaskGroupID
}

// Parent returns the parent resource.
func (x ParsedTaskGroupName) Parent() ParsedJobName {
	return ParsedJobName{
		ProjectID:  x.ProjectID,
		LocationID: x.LocationID,
		JobID:      x.JobID,
	}
}

// String returns the string representation.
func (x ParsedTaskGroupName) String() string {
	return x.Name()
}

// Key returns the key of the resource.
func (x ParsedTaskName) Key() string {
	return x.TaskID
}

// Parent returns the parent resource.
func (x ParsedTaskName) Parent() ParsedTaskGroupName {
	return ParsedTaskGroupName{
		ProjectID:   x.ProjectID,
		LocationID:  x.LocationID,
		JobID:       x.JobID,
		TaskGroupID: x.TaskGroupID,
	}
}

// String returns the string representation.
func (x ParsedTaskName) String() string {
	return x.Name()
}
