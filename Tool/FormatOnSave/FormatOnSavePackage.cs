﻿using EnvDTE;
using Microsoft;
using Microsoft.VisualStudio.Shell;
using Microsoft.VisualStudio.Shell.Interop;
using System;
using System.IO;
using System.Text;
using System.Runtime.InteropServices;
using System.Threading;
using Task = System.Threading.Tasks.Task;

namespace FormatOnSave
{
	/// <summary>
	/// This is the class that implements the package exposed by this assembly.
	/// </summary>
	/// <remarks>
	/// <para>
	/// The minimum requirement for a class to be considered a valid package for Visual Studio
	/// is to implement the IVsPackage interface and register itself with the shell.
	/// This package uses the helper classes defined inside the Managed Package Framework (MPF)
	/// to do it: it derives from the Package class that provides the implementation of the
	/// IVsPackage interface and uses the registration attributes defined in the framework to
	/// register itself and its components with the shell. These attributes tell the pkgdef creation
	/// utility what data to put into .pkgdef file.
	/// </para>
	/// <para>
	/// To get loaded into VS, the package must be referred by &lt;Asset Type="Microsoft.VisualStudio.VsPackage" ...&gt; in .vsixmanifest file.
	/// </para>
	/// </remarks>
	[PackageRegistration(UseManagedResourcesOnly = true, AllowsBackgroundLoading = true)]
	[Guid(FormatOnSavePackage.PackageGuidString)]
	[ProvideAutoLoad(UIContextGuids80.SolutionExists, PackageAutoLoadFlags.BackgroundLoad)]
	public sealed class FormatOnSavePackage : AsyncPackage
	{
		/// <summary>
		/// FormatOnSavePackage GUID string.
		/// </summary>
		public const string PackageGuidString = "02cd9be2-1281-45eb-96b5-1fc5bbf940d2";

		#region Package Members

		/// <summary>
		/// Initialization of the package; this method is called right after the package is sited, so this is the place
		/// where you can put all the initialization code that rely on services provided by VisualStudio.
		/// </summary>
		/// <param name="cancellationToken">A cancellation token to monitor for initialization cancellation, which can occur when VS is shutting down.</param>
		/// <param name="progress">A provider for progress updates.</param>
		/// <returns>A task representing the async work of package initialization, or an already completed task if there is none. Do not return null from this method.</returns>
		protected override async Task InitializeAsync(CancellationToken cancellationToken, IProgress<ServiceProgressData> progress)
		{
			// When initialized asynchronously, the current thread may be a background thread at this point.
			// Do any initialization that requires the UI thread after switching to the UI thread.
			await this.JoinableTaskFactory.SwitchToMainThreadAsync(cancellationToken);

			m_dte = await GetServiceAsync(typeof(DTE)) as DTE;
			Assumes.Present(m_dte);
			m_documentEvents = m_dte.Events.DocumentEvents;
			m_documentEvents.DocumentSaved += DocumentEvents_DocumentSaved;
		}

		private void DocumentEvents_DocumentSaved(Document document)
		{
			try
			{
				ThreadHelper.ThrowIfNotOnUIThread();

				m_dte.ExecuteCommand("Edit.FormatDocument");

				var stream = new FileStream(document.FullName, FileMode.Open);
				var reader = new StreamReader(stream, Encoding.Default, true);
				reader.Read();

				if (reader.CurrentEncoding.EncodingName == Encoding.UTF8.EncodingName &&
					reader.CurrentEncoding.GetPreamble().Length != 0)
				{
					stream.Close();
				}
				else
				{
					try
					{
						stream.Position = 0;
						reader = new StreamReader(stream, new UTF8Encoding(false, true));
						string text = reader.ReadToEnd();
						stream.Close();
						File.WriteAllText(document.FullName, text, new UTF8Encoding(false));
					}
					catch (DecoderFallbackException)
					{
						stream.Position = 0;
						reader = new StreamReader(stream, Encoding.Default);
						string text = reader.ReadToEnd();
						stream.Close();
						File.WriteAllText(document.FullName, text, new UTF8Encoding(false));
					}
				}
			}
			catch (Exception e)
			{
				var outputWindow = Package.GetGlobalService(typeof(SVsOutputWindow)) as IVsOutputWindow;

				var paneGuid = Microsoft.VisualStudio.VSConstants.OutputWindowPaneGuid.GeneralPane_guid;
				outputWindow.CreatePane(paneGuid, "Mine!", 1, 0);
				outputWindow.GetPane(paneGuid, out IVsOutputWindowPane pane);

				pane.OutputString($"Exception occured. [document: {document.FullName}, exception: {e}]");
			}
		}

		private DTE m_dte;
		private DocumentEvents m_documentEvents;

		#endregion
	}
}