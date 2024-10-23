package types

type Descriptor interface{}

/*
 * Title: Title of this group of files.
 * Alias: The alias of the file or directory.
 * GroupID: The unique id of this group.
 * Details: The description of this group of file.
 * Tags: The tags of this group of files.
 * ChildFiles: File ids in thisfile group, not Sid.
 */
type FGroup struct {
	Title string
	Alias []string
	GroupID int
	Details string
	Tags []string
	ChildFiles []int
}

/*
 * FDescriptor is a struct to store the file information, including dir and file
 * Public: DescriptorP struct to store the public information.
 * Path: the abs path of the file, hidden to public.
 */
type FDescriptor struct {
	Path  string      `json:"path"`
	Public DescriptorP `json:"public"`
}

/*
 * DescriptorP is a struct to store the public information of the file or directory
 * Ext: The ext name of the file
		 Folders usually leave this field empty
 * FullName: The full name of the file or directory
			  Usually, only file has this field, this includes file name and ext name
			  Folders' Fullname usually is the Name, so just leave FullName empty
 * ID: The id of the file or directory
 * Sid: If the file or dir is part of a file group, it has a unique sid in this file group.
        Format: "<series>:<episode>", such as: "1:12", "0:6.5" or "2:" (A folder).
        If it is a top folder of a file group, set it to ":<GroupID>"
 * Name: The name of the file or directory
 * Parent: The parent folder id.
 * Related: The related files of the file
 * SubFiles: the id of the files in the directory
*/
type DescriptorP struct {
	Ext      string   `json:"ext"`
	FullName string   `json:"fullName"`
	Hidden   bool     `json:"hidden"`
	ID       int64    `json:"id"`
	Sid	string	`json:"sid"`
	IsDir    bool     `json:"isDir"`
	Name     string   `json:"name"`
	Parent   int64   `json:"parent"`
	Related  []string `json:"related"`
	SubFiles []int    `json:"subFiles"`
}
