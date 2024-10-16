package types

type Descriptor interface{}

/*
 * FileDescriptor is a struct to store the file information

 * Private: DescriptorP struct to store the private information

 * Path: the abs path of the file
 */

type FileDescriptor struct {
	Path    string      `json:"path"`
	Private DescriptorP `json:"private"`
}

/*
 * DirDescriptor is a struct to store the directory information

 * Private: DescriptorP struct to store the private information

 * Path: the abs path of the directory
 */

type DirDescriptor struct {
	Path    string      `json:"path"`
	Private DescriptorP `json:"private"`
}

/*
 * DescriptorP is a struct to store the private information of the file or directory

 * Alias: The alias of the file or directory
 *         Usually, only the dir has alias, sub files share the same alias

 * Details: The description of the file or directory
			 Like Tags, usually, only the dir has details, sub files share the same details

 * Ext: The ext name of the file
		 Folders usually leave this field empty

 * FullName: The full name of the file or directory
			  Usually, only file has this field, this includes file name and ext name
			  Folders' Fullname usually is the Name, so just leave FullName empty

 * ID: The id of the file or directory

 * Name: The name of the file or directory

 * Parent: id of the parent folder,
			WatchPath is 0, so the descriptor whose parent is 0 will be exported first.

 * Related: The related files of the file

 * SubFiles: the id of the files in the directory

 * Tags: The tags of the file or directory
          Usually, only the dir has tags, sub files share the same tags.
		  If the file is a single file, the tags should be set.
		  But it is advised to store the file in a separate folder and set the tags in the folder.
*/

type DescriptorP struct {
	Alias    []string `json:"alias"`
	Details  string   `json:"details"`
	Ext      string   `json:"ext"`
	FullName string   `json:"fullName"`
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Parent   int      `json:"parent"`
	Related  []string `json:"related"`
	SubFiles []int    `json:"subFiles"`
	Tags     []string `json:"tags"`
}
